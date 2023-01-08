package transactiondto

type CreateTransaction struct {
	// ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int    `json:"qty"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	TripID     int    `json:"trip_id"`
	UserID     int    `json:"user_id"`
}
type UpdateTransaction struct {
	// ID         int    `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int    `json:"qty"`
	Total      int    `json:"total"`
	Status     string `json:"status"`
	Attachment string `json:"attachment"`
	TripID     int    `json:"trip_id"`
	UserID     int    `json:"user_id"`
}
