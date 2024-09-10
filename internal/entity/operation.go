package entity

import (
	"errors"
	"time"
)

type Operation struct {
	ID      int       `json:"id"`
	Amount  int       `json:"amount"`
	Type    string    `json:"type"` //topup or deduct
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}

// вместо переменных в var по сути нужно использовать что-то подобное
type ResponseData struct {
	Account int `json:"account"`
	//Nums_of_operations int               `json:"nums_of_operations"`
	Operations map[int]Operation `json:"operations"`
}

// иммитация бд
var (
	Account            = 0
	Nums_of_operations = 10
	Operations         = map[int]Operation{
		1: {1, 200, "topup", "add 200", time.Date(2020, time.April,
			11, 21, 34, 01, 0, time.UTC)},
		2: {2, 200, "topup", "add more 200", time.Date(2020, time.April,
			12, 21, 34, 01, 0, time.UTC)},
		3: {3, 500, "deduct", "", time.Date(2020, time.April,
			13, 21, 34, 01, 0, time.UTC)},
		4: {4, 250, "deduct", "oops decrease", time.Date(2020, time.April,
			14, 21, 34, 01, 0, time.UTC)},
		5: {5, 250, "topup", "oops topup", time.Date(2020, time.April,
			15, 21, 34, 01, 0, time.UTC)},
		6: {6, 250, "topup", "oops topup", time.Date(2020, time.April,
			16, 21, 34, 01, 0, time.UTC)},
		7: {7, 250, "deduct", "oops decrease", time.Date(2020, time.April,
			17, 21, 34, 01, 0, time.UTC)},
		8: {8, 250, "deduct", "", time.Date(2020, time.April,
			18, 21, 34, 01, 0, time.UTC)},
		9: {9, 250, "topup", "oops topup", time.Date(2020, time.April,
			19, 21, 34, 01, 0, time.UTC)},
		10: {10, 250, "deduct", "oops decrease", time.Date(2020, time.April,
			20, 21, 34, 01, 0, time.UTC)},
	}
)

// не знаю, куда это
func MakeOperation(operation Operation) error {
	var val int
	if operation.Type == "deduct" {
		val = Account - operation.Amount
	} else {
		val = Account + operation.Amount
	}
	if val < 0 {
		return errors.New("Текущий счёт низкий, операция невозможна")
	}
	Nums_of_operations += 1
	Operations[Nums_of_operations] = Operation{Nums_of_operations, operation.Amount, operation.Type, operation.Comment, operation.Date}
	Account = val
	return nil
}

func GetStart() int {
	start := Nums_of_operations - 10
	if start < 0 {
		start = 0
	}
	return start
}
