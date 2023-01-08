package authdto

type RegisterResponse struct {
	FullName string `json:"fullname"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`

	Token string `gorm:"type: varchar(255)" json:"token"`
}
type LoginResponse struct {
	FullName string `json:"fullname"`
	Email    string `json:"email" form:"email" validate:"required"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Role     string `json:"role"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
}
