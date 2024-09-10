package controller

import (
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"project/pkg/logger"
)

// потом сюда еще бд
type Server struct {
	Router *mux.Router
}

func New(webDir string) *Server {
	s := &Server{mux.NewRouter()}
	s.Router.HandleFunc("/", s.HomeHandler)
	s.Router.HandleFunc("/topup", s.topupHandler).Methods("PUT")
	s.Router.HandleFunc("/deduct", s.deductHandler).Methods("PUT")
	s.Router.HandleFunc("/list", s.listHandler).Methods("GET")

	//скорее всего надо таким образом реализовывать /, но не до конца работает
	//s.Router.PathPrefix("/").Handler(http.FileServer(http.Dir(webDir)))
	return s

}

func (s *Server) Run(addr string) {
	logger.Logger.Info("Сервер запущен на http://127.0.0.1:8080")
	if err := http.ListenAndServe("localhost"+addr, s.Router); err != nil {
		logger.Logger.Error("Fatal Error", slog.Int("status", http.StatusBadGateway))
	}
}
