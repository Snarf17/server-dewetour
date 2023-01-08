package models

type User struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname" gorm:"type: varchar(255)"`
	Email    string `json:"email" gorm:"type: varchar(255)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Phone    string `json:"phone" `
	Address  string `json:"address" gorm:"type: text"`
	Role     string `json:"role"`
}

type UsersResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

func (UsersResponse) TableName() string {
	return "users"
}
