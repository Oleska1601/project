package usecase

import (
	"fmt"
	"project/internal/entity"
	"project/internal/usecase/repo/sqlitedb"
)

// сохраняем измененные данные в ипровизированную бд
func MakeOperation(db *sqlitedb.SqliteDB, operation entity.Operation) error {
	var val int
	var account entity.Account

	if operation.Type == "deduct" {
		val = account.Balance - operation.Amount
	} else {
		val = account.Balance + operation.Amount
	}
	if val < 0 {
		return fmt.Errorf("Текущий счёт низкий, операция невозможна")
	}
	operation.ID += 1

	if err := db.InsertTableOperations(operation); err != nil {
		return fmt.Errorf("fail to insert operation in db")
	}
	if err := db.UpdateTableAccounts(account); err != nil {
		return fmt.Errorf("fail to update account in db")
	}
	return nil
}

/*
	data, err := ioutil.ReadFile("internal/usecase/repo.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return err
	}

*/
/*
	updateData, _ := json.Marshal(responseData)
	err = ioutil.WriteFile("internal/usecase/repo.json", updateData, 0777)
	if err != nil {
		return err
	}
	return nil
*/
