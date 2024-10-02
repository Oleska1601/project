package controller

import (
	"encoding/json"
	"io/ioutil"
	"log/slog"
	"net/http"
	"project/internal/entity"
	"project/internal/usecase"
	"project/pkg/logger"
)

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	/*
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
	*/
	var responseData entity.ResponseData
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&responseData); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		logger.Logger.Error("Bad Request", slog.Int("status", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	updateData, _ := json.Marshal(responseData)
	//запись в бд
	ioutil.WriteFile("internal/usecase/repo.json", updateData, 0644)
	w.WriteHeader(http.StatusCreated)
	logger.Logger.Info("Content is created", slog.Int("status", http.StatusOK))

}

func (s *Server) topupHandler(w http.ResponseWriter, r *http.Request) {
	//put
	var operation entity.Operation
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		logger.Logger.Error("Bad Request", slog.Int("status", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	err := usecase.MakeOperation(operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error("Impossible to make operation", slog.Int("status", http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	logger.Logger.Info("Operation topup is made", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(operation)

}

func (s *Server) deductHandler(w http.ResponseWriter, r *http.Request) {
	//put
	var operation entity.Operation
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		logger.Logger.Error("Bad Request", slog.Int("status", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	err := usecase.MakeOperation(operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error("Impossible to make operation", slog.Int("status", http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	logger.Logger.Info("Operation deduct is made", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(operation)
}

// get
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	var responseData entity.ResponseData
	w.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadFile("internal/usecase/repo.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error("Cannot find database", slog.Int("status", http.StatusInternalServerError))
		return
	}
	if err = json.Unmarshal(data, &responseData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error("Impossible to read data", slog.Int("status", http.StatusInternalServerError))
		return
	}
	responseData.Operations = usecase.GetLastTenOperations(responseData)
	w.WriteHeader(http.StatusOK)
	logger.Logger.Info("Get last 10 made operations", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(responseData.Operations)
}
