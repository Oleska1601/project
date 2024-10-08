package entity

type Operation struct {
	ID        int    `json:"id"`
	AccountID int    `json:"account_id"`
	Amount    int    `json:"amount"`
	Type      string `json:"type"` //topup or deduct
	Comment   string `json:"comment"`
	Date      string `json:"date"`
}
