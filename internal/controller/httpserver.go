package controller

import (
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"log/slog"
	"net/http"
	"project/config"
	"project/internal/usecase"
	"project/pkg/logger"
)

type Server struct {
	Router *mux.Router
	u      *usecase.Usecase
	log    *logger.Logger
}

func New(cfg *config.Config, u *usecase.Usecase, log *logger.Logger) *Server {
	s := &Server{mux.NewRouter(), u, log}
	//webDir := cfg.Web.Path
	s.Router.HandleFunc("/", s.HomeHandler).Methods("POST")
	s.Router.HandleFunc("/topup", s.topupHandler).Methods("PUT")
	s.Router.HandleFunc("/deduct", s.deductHandler).Methods("PUT")
	s.Router.HandleFunc("/list", s.listHandler).Methods("GET")

	s.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return s

}

func (s *Server) Run(cfg *config.Config) {
	s.log.Info("Сервер запущен на http://127.0.0.1:" + cfg.Port)
	if err := http.ListenAndServe("localhost:"+cfg.Port, s.Router); err != nil {
		s.log.Error(err, slog.String("msg", "fatal error"), slog.Int("status", http.StatusBadGateway))
	}
}
