package pg

import (
	"database/sql"
	"fmt"
	"gophermart/internal/config"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"

	_ "github.com/lib/pq"
)

type PGDatabase struct {
	db     *sql.DB
	logger *logger.Logger
}

func NewPGDatabase(config *config.Config, logger *logger.Logger) repository.DatabaseRepository {
	connStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%s",
		config.DBName,
		config.User,
		config.Password,
		config.Host,
		config.Port,
	)
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
