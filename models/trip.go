package models

import "time"

type Trip struct {
	ID             int             `json:"id" gorm:"primary_key:auto_increment"`
	Title          string          `json:"title" gorm:"type: varchar(255)"`
	CountryID      int             `json:"-" `
	Country        CountryResponse `json:"country" `
	Accomodation   string          `json:"accomodation" gorm:"type: varchar(255)"`
	Transportation string          `json:"transportation" gorm:"type: varchar(255)"`
	Eat            string          `json:"eat" gorm:"type: varchar(255)"`
	Day            int             `json:"day" gorm:"type: int"`
	Night          int             `json:"night" gorm:"type: int"`
	DateTrip       time.Time       `json:"date_trip" form:"date_trip"`
	Price          int             `json:"price" gorm:"type: int"`
	Quota          int             `json:"quota" gorm:"type: int"`
	Description    string          `json:"description" gorm:"type: varchar(255)"`
	Image          string          `json:"image" gorm:"type: varchar(255)"`
	UserID         int             `json:"user_id"`
	User           UsersResponse   `json:"user"`
}

type TripsResponse struct {
	ID             int             `json:"id"`
	Title          string          `json:"title"`
	CountryID      int             `json:"country_id"`
	Country        CountryResponse `json:"country"`
	Accomodation   string          `json:"accomodation"`
	Transportation string          `json:"transportation"`
	Eat            string          `json:"eat"`
	Day            int             `json:"day"`
	Night          int             `json:"night"`
	DateTrip       time.Time       `json:"date_trip" form:"date_trip"`
	Price          int             `json:"price"`
	Quota          int             `json:"quota"`
	Description    string          `json:"description"`
	Image          string          `json:"image"`
}

func (TripsResponse) TableName() string {
	return "trips"
}
