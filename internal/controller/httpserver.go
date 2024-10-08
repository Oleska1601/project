package controller

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"project/config"
	"project/internal/usecase/repo/sqlitedb"
	"project/pkg/logger"
)

type Server struct {
	Router *mux.Router
	db     *sqlitedb.SqliteDB
}

func New(db *sqlitedb.SqliteDB, filename string) *Server {

	s := &Server{mux.NewRouter(), db}
	s.Router.HandleFunc("/", s.HomeHandler).Methods("POST")
	s.Router.HandleFunc("/topup", s.topupHandler).Methods("PUT")
	s.Router.HandleFunc("/deduct", s.deductHandler).Methods("PUT")
	s.Router.HandleFunc("/list", s.listHandler).Methods("GET")

	//subRouter := s.Router.PathPrefix("/api").Subrouter() //вернет new router
	return s

}

func (s *Server) Run(cfg *config.Config) {
	logger.Logger.Info("Сервер запущен на http://127.0.0.1:" + cfg.Port)
	if err := http.ListenAndServe("localhost:"+cfg.Port, s.Router); err != nil {
		logger.Logger.Error("Fatal Error", slog.Int("status", http.StatusBadGateway))
	}
}
