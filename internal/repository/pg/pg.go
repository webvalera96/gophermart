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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	_ "github.com/lib/pq"
)

type PGDatabase struct {
	db     *sql.DB
	logger *logger.Logger
}

//go:embed migrations
var embedMigrations embed.FS

func migrateDatabase(config *config.Config, logger *logger.Logger) error {

	d, err := iofs.New(embedMigrations, "migrations")
	if err != nil {
		logger.Fatalf("Migration failed: %v", err)
		return err
	}

	m, err := migrate.NewWithSourceInstance(
		"iofs",
		d,
		//TODO: add configuration for sslmode
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DBName),
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

func NewPGDatabase(config *config.Config, logger *logger.Logger) repository.DatabaseRepository {
	// TODO: add configuration for sslmode
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s sslmode=disable",
		config.DBName,
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
	err := migrateDatabase(config, logger)

	db, err := sql.Open("postgres", connStr)
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
	pg.logger.Debug("TODO: create user in database")
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

func (pg PGDatabase) CreateOrder(order models.Order) (*models.Order, error) {
	var savedOrder models.Order

	user, err := pg.GetUserByLogin(order.Login)

	var unf *repository.UserNotFoundError
	if errors.As(err, &unf) {
		return nil, &repository.UserNotFoundError{Msg: "user not found"}
	}
	var orderId int
	var orderDate string
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
	var orderDate string
	row := pg.db.QueryRow("SELECT id, number, user_id, order_date FROM orders WHERE number = $1", number)
	err := row.Scan(&id, &order.Number, &userId, &orderDate)
	order.SetID(id)
	order.SetOrderDate(orderDate)

	if err == sql.ErrNoRows { // не удалось найти заказ
		return nil, &repository.OrderNotFoundError{}
	} else if err != nil {
		return nil, err
	}

	return &order, nil
}
