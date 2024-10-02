package usecase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"project/internal/entity"
)

// сохраняем измененные данные в ипровизированную бд
func MakeOperation(operation entity.Operation) error {
	var val int
	var responseData entity.ResponseData
	data, err := ioutil.ReadFile("internal/usecase/repo.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &responseData)
	if err != nil {
		return err
	}
	if operation.Type == "deduct" {
		val = responseData.Account - operation.Amount
	} else {
		val = responseData.Account + operation.Amount
	}
	if val < 0 {
		return fmt.Errorf("Текущий счёт низкий, операция невозможна")
	}
	operation.ID += 1
	responseData.NumsOfOperations += 1
	responseData.Operations = append(responseData.Operations, operation)
	responseData.Account = val
	updateData, _ := json.Marshal(responseData)
	err = ioutil.WriteFile("internal/usecase/repo.json", updateData, 0777)
	if err != nil {
		return err
	}
	return nil
}
