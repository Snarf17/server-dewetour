package tripdto

import "time"

type CreateTripRequest struct {
	// ID              int       `json:`
	Title          string    `json:"title" form:"title"`
	CountryID      int       `json:"country_id" form:"country_id"`
	Accomodation   string    `json:"accomodation" form:"accomodation"`
	Transportation string    `json:"transport" form:"transport"`
	Eat            string    `json:"eat" form:"eat"`
	Day            int       `json:"day" form:"day"`
	Night          int       `json:"night" form:"night"`
	DateTrip       time.Time `json:"date_trip" form:"date_trip"`
	Price          int       `json:"price" form:"price"`
	Quota          int       `json:"quota" form:"quota"`
	Description    string    `json:"desc" form:"desc"`
	Image          string    `json:"image" form:"image"`
}

type UpdateTripRequest struct {
	// ID              int       `json:`
	Title          string    `json:"title" form:"title"`
	CountryID      int       `json:"country_id" form:"country_id"`
	Accomodation   string    `json:"accomodation" form:"accomodation"`
	Transportation string    `json:"transport" form:"transport"`
	Eat            string    `json:"eat" form:"eat"`
	Day            int       `json:"day" form:"day"`
	Night          int       `json:"night" form:"night"`
	DateTrip       time.Time `json:"date_trip" form:"date_trip"`
	Price          int       `json:"price" form:"price"`
	Quota          int       `json:"quota" form:"quota"`
	Description    string    `json:"desc" form:"desc"`
	Image          string    `json:"image" form:"image"`
	UserID         int       `json:"-"`
}

// type UpdateUserRequest struct {
// 	// ID int `json:`
// 	FullName string `json:"fullname" form:"name" validate:"required"`
// 	Email    string `json:"email" form:"email" validate:"required"`
// 	Password string `json:"password" form:"password" validate:"required"`
// 	Phone    int    `json:"phone" gorm:"type: int(100)" validate:"required"`
// 	Address  string `json:"address" gorm:"type: text" validate:"required"`
// }
