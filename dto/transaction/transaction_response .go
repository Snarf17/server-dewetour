package transactiondto

type TransactionResponse struct {
	// ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int    `json:"qty"`
	Total      int    `json:"total"`
	Attachment string `json:"attachment"`
}
