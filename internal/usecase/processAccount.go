package usecase

import (
	"project/internal/entity"
	"project/internal/usecase/repo/sqlitedb"
)

// получение данных из бд
func GetLastTenOperations(db *sqlitedb.SqliteDB, account entity.Account) ([]entity.Operation, error) {
	operations, err := db.SelectTableOperations(account)
	if err != nil {
		return nil, err
	}
	start := len(operations) - 10
	if start < 0 {
		start = 0
	}
	return operations[start:], nil
}
