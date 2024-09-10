package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"project/internal/entity"
	"project/pkg/logger"
)

//1: как обновлять значение response := entity.ResponseData{entity.Account}, но  чтобы не делать в каждой ф-ии обработчке tpl.Execute(w, response)
//2: как огранизовать работу сразу с двумя структурами, т.е. каким образом в listHandler например понять, что json.NewDecoder(r.Body).Decode(&response) будет принимать параметры ResponseData, а не Operation
//3: если указывать метод post  s.Router.HandleFunc("/", s.HomeHandler).Methods("POST"), то перестает работать, поскольку в homeHandler я ничего не считываю и не передаю обратно пользователю -> вопрос: как адекватно оформить HomeHandler
//4: как нормально разбить css, js по файлам, тк они не обрабатываются, если их указывать только в html

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	//чтение файлов шаблона
	tpl, err := template.ParseFiles("front/index.html")
	if err != nil {
		logger.Logger.Error("Internal Server Error", slog.Int("status", http.StatusInternalServerError))
		return
	}
	response := entity.ResponseData{entity.Account, entity.Operations}
	// Затем мы используем метод Execute() для записи содержимого
	// шаблона в тело HTTP ответа. Последний параметр в Execute() предоставляет
	// возможность отправки динамических данных в шаблон.
	err = tpl.Execute(w, response)
	if err != nil {
		logger.Logger.Error("Internal Server Error", slog.Int("status", http.StatusInternalServerError))
		return
	}
}

func (s *Server) topupHandler(w http.ResponseWriter, r *http.Request) {
	//put
	var operation entity.Operation
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		logger.Logger.Error("Bad Request", slog.Int("status", 400))
		return
	}
	defer r.Body.Close()
	err := entity.MakeOperation(operation)
	if err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		logger.Logger.Error("Bad Request", slog.Int("status", 400))
		return
	}

	w.WriteHeader(http.StatusOK)
	logger.Logger.Info("Operation topup is made", slog.Int("status", 200))
	fmt.Fprintln(w, "текущий счет:", entity.Account)
	json.NewEncoder(w).Encode(operation)

}

func (s *Server) deductHandler(w http.ResponseWriter, r *http.Request) {
	//put
	var operation entity.Operation
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		logger.Logger.Error("Bad Request", slog.Int("status", 400))
		return
	}
	defer r.Body.Close()
	err := entity.MakeOperation(operation)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "текущий счет:", entity.Account)
	logger.Logger.Info("Operation deduct is made", slog.Int("status", 200))
	json.NewEncoder(w).Encode(operation)
}

// кривая реализация, надо, скорее всего, переделать, как-то используя json.NewDecoder(r.Body).Decode(&response), где response = entity.ResponseData
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	start := entity.GetStart()
	for i := start; i < i+10; i++ {
		fmt.Fprintln(w, entity.Operations[i])
	}
}
