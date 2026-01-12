// Package pg предоставляет реализацию DatabaseRepository для PostgreSQL.
package pg

import (
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"gophermart/internal/config"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

// PGDatabase представляет реализацию DatabaseRepository для PostgreSQL.
// Содержит подключение к базе данных и логгер.
type PGDatabase struct {
	db     *sql.DB
	logger *logger.Logger
}

//go:embed migrations
var embedMigrations embed.FS

func migrateDatabase(config *config.Config, logger *logger.Logger) error {
	if config.DatabaseURI == "" {
		logger.Fatalf("DatabaseURI is not set")
		return fmt.Errorf("DatabaseURI is not set")
	}

	d, err := iofs.New(embedMigrations, "migrations")
	if err != nil {
		logger.Fatalf("Migration failed: %v", err)
		return err
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		d,
		config.DatabaseURI,
	)
	if err != nil {
		logger.Fatalf("Migration failed: %v", err)
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		logger.Fatalf("Migration failed: %v", err)
	}

	return nil
}

// NewPGDatabase создает новый экземпляр PGDatabase и выполняет миграции базы данных.
// Принимает config - конфигурация приложения с параметрами подключения к БД,
// logger - логгер для записи сообщений.
// Возвращает реализацию DatabaseRepository для PostgreSQL.
// Паникует, если не удалось подключиться к базе данных или выполнить миграции.
func NewPGDatabase(config *config.Config, logger *logger.Logger) repository.DatabaseRepository {
	if config.DatabaseURI == "" {
		logger.Fatalf("DatabaseURI is not set")
		panic("DatabaseURI is not set")
	}

	err := migrateDatabase(config, logger)
	if err != nil {
		logger.Fatalf("Migration failed: %v", err)
		panic(err)
	}

	// Преобразуем postgres:// URI в формат для lib/pq
	// lib/pq использует формат "postgres://user:password@host:port/dbname?sslmode=disable"
	// который совместим с нашим DatabaseURI
	db, err := sql.Open("postgres", config.DatabaseURI)
	if err != nil {
		panic(err)
	}
	return PGDatabase{db: db, logger: logger}
}

// Create user in postgresql database
func (pg PGDatabase) CreateUser(u models.User) error {
	var id int
	_, err := pg.GetUserByLogin(u.Login)
	var unf *repository.UserNotFoundError
	if errors.As(err, &unf) { // проверяем, что пользователь с таким именем не будет зарегестрирован
		err := pg.db.QueryRow("INSERT INTO users(login, password) VALUES ($1, $2) RETURNING id", u.Login, u.Password).Scan(&id)
		if err != nil {
			return err
		}
	} else { // если такой пользователь уже существует, вернуть ошибку создания пользователя
		return &repository.UserAlreadyExistsError{}
	}
	return nil
}

func (pg PGDatabase) GetUserByLogin(login string) (*models.User, error) {
	var user models.User
	var id int
	row := pg.db.QueryRow("SELECT id, login, password FROM users WHERE login = $1", login)
	err := row.Scan(&id, &user.Login, &user.Password)
	user.SetID(id)

	if err == sql.ErrNoRows { // не удалось найти пользователя
		return nil, &repository.UserNotFoundError{}
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (pg PGDatabase) GetUserBalance(login string) (*models.User, error) {
	var user models.User
	var id int
	var currentBalance, withdrawn sql.NullFloat64

	row := pg.db.QueryRow(
		"SELECT id, login, COALESCE(current_balance, 0), COALESCE(withdrawn, 0) FROM users WHERE login = $1",
		login,
	)
	err := row.Scan(&id, &user.Login, &currentBalance, &withdrawn)
	user.SetID(id)

	if err == sql.ErrNoRows {
		return nil, &repository.UserNotFoundError{}
	} else if err != nil {
		return nil, err
	}

	if currentBalance.Valid {
		user.SetCurrentBalance(currentBalance.Float64)
	} else {
		user.SetCurrentBalance(0.0)
	}

	if withdrawn.Valid {
		user.SetWithdrawn(withdrawn.Float64)
	} else {
		user.SetWithdrawn(0.0)
	}

	return &user, nil
}

func (pg PGDatabase) WithdrawBalance(login string, orderNumber string, sum float64) error {
	// Получаем текущий баланс пользователя
	user, err := pg.GetUserBalance(login)
	if err != nil {
		return err
	}

	// Проверяем, что на счету достаточно средств
	if user.GetCurrentBalance() < sum {
		return &repository.InsufficientFundsError{}
	}

	// Начинаем транзакцию
	tx, err := pg.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Обновляем баланс: уменьшаем current_balance и увеличиваем withdrawn
	_, err = tx.Exec(
		"UPDATE users SET current_balance = current_balance - $1, withdrawn = withdrawn + $1 WHERE login = $2",
		sum,
		login,
	)
	if err != nil {
		return err
	}

	// Сохраняем запись о списании
	_, err = tx.Exec(
		"INSERT INTO withdrawals(user_id, order_number, sum) VALUES ($1, $2, $3)",
		user.GetID(),
		orderNumber,
		sum,
	)
	if err != nil {
		return err
	}

	// Коммитим транзакцию
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (pg PGDatabase) CreateOrder(order models.Order) (*models.Order, error) {
	var savedOrder models.Order

	user, err := pg.GetUserByLogin(order.Login)

	var unf *repository.UserNotFoundError
	if errors.As(err, &unf) {
		return nil, &repository.UserNotFoundError{Msg: "user not found"}
	}
	var orderId int
	var orderDate time.Time
	err = pg.db.QueryRow("INSERT INTO orders(number, user_id) VALUES ($1, $2) RETURNING id, order_date",
		order.Number,
		user.GetID(),
	).Scan(&orderId, &orderDate)

	if err != nil {
		return nil, err
	}

	savedOrder.SetID(orderId)
	savedOrder.SetOrderDate(orderDate)
	savedOrder.Login = order.Login
	savedOrder.Number = order.Number

	return &savedOrder, nil
}

func (pg PGDatabase) GetOrderByNumber(number string) (*models.Order, error) {
	var order models.Order
	var id int
	var userId int
	var orderDate time.Time
	var status sql.NullString
	var accrual sql.NullFloat64
	var login sql.NullString

	row := pg.db.QueryRow(
		"SELECT o.id, o.number, o.user_id, o.order_date, COALESCE(o.status, 'REGISTERED'), o.accrual, u.login FROM orders o JOIN users u ON o.user_id = u.id WHERE o.number = $1",
		number,
	)
	err := row.Scan(&id, &order.Number, &userId, &orderDate, &status, &accrual, &login)
	order.SetID(id)
	order.SetOrderDate(orderDate)

	if err == sql.ErrNoRows { // не удалось найти заказ
		return nil, &repository.OrderNotFoundError{}
	} else if err != nil {
		return nil, err
	}

	if status.Valid {
		order.SetStatus(status.String)
	} else {
		order.SetStatus("REGISTERED")
	}

	if accrual.Valid {
		accrualValue := accrual.Float64
		order.SetAccrual(&accrualValue)
	} else {
		order.SetAccrual(nil)
	}

	if login.Valid {
		order.Login = login.String
	}

	return &order, nil
}

func (pg PGDatabase) GetOrdersByUserLogin(login string) ([]models.Order, error) {
	var orders []models.Order

	rows, err := pg.db.Query(
		"SELECT o.id, o.number, o.order_date FROM orders o JOIN users u ON o.user_id = u.id WHERE u.login = $1",
		login,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var order models.Order
		var id int
		var orderDate time.Time
		err := rows.Scan(&id, &order.Number, &orderDate)
		if err != nil {
			return nil, err
		}
		err = order.SetID(id)
		if err != nil {
			return nil, err
		}
		err = order.SetOrderDate(orderDate)
		if err != nil {
			return nil, err
		}
		order.Login = login
		orders = append(orders, order)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (pg PGDatabase) GetWithdrawalsByUserLogin(login string) ([]models.Withdrawal, error) {
	var withdrawals []models.Withdrawal

	rows, err := pg.db.Query(
		"SELECT w.id, w.order_number, w.sum, w.processed_at FROM withdrawals w JOIN users u ON w.user_id = u.id WHERE u.login = $1 ORDER BY w.processed_at DESC",
		login,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var withdrawal models.Withdrawal
		var id int
		var processedAt time.Time
		err := rows.Scan(&id, &withdrawal.OrderNumber, &withdrawal.Sum, &processedAt)
		if err != nil {
			return nil, err
		}
		withdrawal.SetID(id)
		withdrawal.SetProcessedAt(processedAt)
		withdrawals = append(withdrawals, withdrawal)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return withdrawals, nil
}
