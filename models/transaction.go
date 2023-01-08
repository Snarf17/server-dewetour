package models

type Transaction struct {
	ID         int           `json:"id" gorm:"primary_key:auto_increment"`
	CounterQty int           `json:"qty"`
	Total      int           `json:"total"`
	Status     string        `json:"status"`
	Attachment string        `json:"attachment"`
	TripID     int           `json:"trip_id"`
	Trip       TripsResponse `json:"trip"`
	UserID     int           `json:"user_id"`
	User       UsersResponse `json:"user"`
	// UserID     int           `json:"-"`
	// User       UserResponse  `json:"user"`
}
