package sqlitedb

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"project/pkg/logger"
)

type SqliteDB struct {
	db *sql.DB
}

func New(filename string) (*SqliteDB, error) {
	db, err := sql.Open("sqlite3", filename)
	if err != nil {
		return nil, err
	}
	database := &SqliteDB{db}
	err = database.CreateTableAccounts()
	if err != nil {
		logger.Logger.Error("impossible to create table in database")
	}
	err = database.CreateTableOperations()
	if err != nil {
		logger.Logger.Error("impossible to create table in database")
	}
	logger.Logger.Info("tables in database are created")
	return database, nil
}

func (database *SqliteDB) Close() error {
	return database.db.Close()
}
