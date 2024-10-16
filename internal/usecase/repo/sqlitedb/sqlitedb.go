package sqlitedb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"project/config"
	"project/pkg/logger"
)

type SqliteDB struct {
	db *sql.DB
}

func New(cfg *config.Config, log *logger.Logger) (*SqliteDB, error) {
	db, err := sql.Open("sqlite3", cfg.DB.Path)
	if err != nil {
		return nil, err
	}
	log.Info("database is opened")
	database := &SqliteDB{db}
	err = database.CreateTableAccounts()
	if err != nil {
		log.Error(err, slog.String("msg", "impossible to create table 'accounts' in database"))
	}
	err = database.CreateTableOperations()
	if err != nil {
		log.Error(err, slog.String("msg", "impossible to create table operations in database"))
	}
	log.Info("tables in database are created")
	return database, nil
}

func (database *SqliteDB) Close(log *logger.Logger) error {
	log.Info("database is closed")
	return database.db.Close()
}
