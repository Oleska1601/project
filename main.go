package main

import (
	_ "github.com/mattn/go-sqlite3"
	"log/slog"
	"project/config"
	"project/internal/controller"
	"project/internal/usecase"
	"project/internal/usecase/repo/sqlitedb"
	"project/pkg/logger"
)

// @title           Swagger Project API
// @version         1.0
// @description     This is a sample for top up or deduct the balance.
// @termsOfService  http://swagger.io/terms/

// @host      localhost:8080
// @BasePath  /

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("Config error: %s", err)
	}
	log := logger.New(cfg)

	db, err := sqlitedb.New(cfg, log)
	if err != nil {
		log.Error(err)
	}
	log.Info("connection to database")
	defer db.Close(log)
	u := usecase.New(db)
	s := controller.New(cfg, u, log)
	s.Run(cfg)
}
