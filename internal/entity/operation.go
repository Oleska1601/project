package entity

import (
	"time"
)

type Operation struct {
	ID      int       `json:"id"`
	Amount  int       `json:"amount"`
	Type    string    `json:"type"` //topup or deduct
	Comment string    `json:"comment"`
	Date    time.Time `json:"date"`
}
