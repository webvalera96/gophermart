package pg

import (
	"database/sql"
	"embed"
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
		//TODO: enable sslmode in production
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.DBName),
	)
	if err != nil {
		logger.Fatalf("Migration failed: %v", err)
		return err
	}
	if err := m.Up(); err != nil {
		logger.Fatalf("Migration failed: %v", err)
	}

	return nil
}

func NewPGDatabase(config *config.Config, logger *logger.Logger) repository.DatabaseRepository {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s",
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
	pg.logger.Debug("TODO: create user in database")
	return nil
}
