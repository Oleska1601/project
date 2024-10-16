package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"project/internal/entity"
	"project/internal/usecase"
)

// HomeHandler базовая страница
// @Summary add account
// @Description add by json account
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body entity.Account true "Add account"
// @Success 200 {object} entity.Account
// @Failure 400 {string} string "Invalid Operation"
// @Router / [post]
func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	var account entity.Account
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервер не может обрабатывать запросы
		s.log.Error(err, slog.String("msg", "Bad Request"), slog.Int("status", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	if err := s.u.DB.InsertTableAccounts(account); err != nil {
		s.log.Error(err, slog.String("msg", "Impossible to insert data into db table accounts"), slog.Int("status", http.StatusBadRequest))
	}
	w.WriteHeader(http.StatusOK)
	s.log.Info("Content is created", slog.Int("status", http.StatusOK))

}

// topupHandler увеличение баланса
// @Summary Topup balance
// @Description topup by json operation with type=topup
// @Tags operations
// @Accept json
// @Produce json
// @Param operation body entity.Operation true "Topup operation"
// @Success 200 {object} entity.Operation
// @Failure 400 {string} string "Invalid Operation"
// @Failure 500 {string} string "Impossible to make operation"
// @Router /topup [put]
func (s *Server) topupHandler(w http.ResponseWriter, r *http.Request) {
	var operation entity.Operation
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		s.log.Error(err, slog.String("msg", "Bad Request"), slog.Int("status", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	err := usecase.MakeOperation(s.u.DB, operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err, slog.String("msg", "Impossible to make operation"), slog.Int("status", http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	s.log.Info("Operation topup is made", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(operation)

}

// deductHandler уменьшение баланса
// @Summary Deduct balance
// @Description deduct by json operation with type=deduct
// @Tags operations
// @Accept json
// @Produce json
// @Param operation body entity.Operation true "Deduct operation"
// @Success 200 {object} entity.Operation
// @Failure 400 {string} string "Invalid Operation"
// @Failure 500 {string} string "Impossible to make operation"
// @Router /deduct [put]
func (s *Server) deductHandler(w http.ResponseWriter, r *http.Request) {
	var operation entity.Operation
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewDecoder(r.Body).Decode(&operation); err != nil {
		http.Error(w, "Invalid Operation", http.StatusBadRequest) //сервее не может обрабатывать запросы
		s.log.Error(err, slog.String("msg", "Bad Request"), slog.Int("status", http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	err := usecase.MakeOperation(s.u.DB, operation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err, slog.String("msg", "Impossible to make operation"), slog.Int("status", http.StatusInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	s.log.Info("Operation deduct is made", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(operation)
}

// listHandler отображение последних 10 операций
// @Summary Get last 10 made operations
// @Description get last 10 made operations
// @Tags operations
// @Accept json
// @Produce json
// @Success 200 {array} entity.Operation
// @Failure 400 {string} string "impossible to get last 10 operations"
// @Router /list [get]
func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	var account entity.Account
	w.Header().Set("Content-Type", "application/json")
	operations, err := usecase.GetLastTenOperations(s.u.DB, account)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		s.log.Error(err, slog.String("msg", "impossible to get last 10 operations"), slog.Int("status", http.StatusInternalServerError))
		return
	}
	w.WriteHeader(http.StatusOK)
	s.log.Info("Get last 10 made operations", slog.Int("status", http.StatusOK))
	json.NewEncoder(w).Encode(operations)
}

//notes
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

/*
	data, err := ioutil.ReadFile("internal/usecase/repo.json")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error("Cannot find database", slog.Int("status", http.StatusInternalServerError))
		return
	}
	if err = json.Unmarshal(data, &account); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logger.Logger.Error("Impossible to read data", slog.Int("status", http.StatusInternalServerError))
		return
	}

*/

//вопросы
//где делать create tables
//где обрабатывать логи из бд и нужно ли (скорее всего нет) http.Error(w, err.Error(), http.StatusInternalServerError)
