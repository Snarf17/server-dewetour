package tripdto

import (
	"time"
)

// type UserResponse struct {
// 	// ID       int    `json:"id"`
// 	FullName string `json:"fullname" form:"name" validate:"required"`
// 	Email    string `json:"email" form:"email" validate:"required"`
// 	Password string `json:"-" form:"name" validate:"required"`
// 	Phone    int    `json:"phone" gorm:"type: int(100)"`
// 	Address  string `json:"address" gorm:"type: text"`
// }

type CountryIdResponse struct {
	Title string `json:"title" form:"title" validate:"required"`
	// CountryID      int       `json:"country_id" form:"country_id" `
	Country        CountryResponse `json:"country_id"`
	Accomodation   string          `json:"accomodation" form:"accomodation" `
	Transportation string          `json:"transport" form:"transport" `
	Eat            string          `json:"eat" form:"eat" `
	Day            int             `json:"day" form:"day" `
	Night          int             `json:"night" form:"night" `
	DateTrip       time.Time       `json:"date_trip" form:"date_trip" `
	Price          int             `json:"price" form:"price" `
	Quota          int             `json:"quota" form:"quota" `
	Description    string          `json:"desc" form:"desc" `
	Image          string          `json:"image" form:"image" `
}

type CountryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (CountryResponse) TableName() string {
	return "countries"
}
