package main

import (
	_ "github.com/mattn/go-sqlite3"
	"project/config"
	"project/internal/controller"
	"project/internal/usecase/repo/sqlitedb"
	"project/pkg/logger"
)

func main() {
	cfg, err := config.New()

	if err != nil {
		logger.Logger.Error(err.Error())
	}
	db, err2 := sqlitedb.New("internal/usecase/repo/sqlitedb/operations.db")
	if err2 != nil {
		logger.Logger.Error(err2.Error())
	}
	logger.Logger.Info("connection to database")
	defer db.Close()
	webDir := "./front"
	s := controller.New(db, webDir)
	s.Run(cfg)
}
