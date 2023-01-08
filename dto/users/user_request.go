package userdto

type CreateUserRequest struct {
	// ID int `json:`
	FullName string `json:"fullname" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" gorm:"type: text" validate:"required"`
}
type UpdateUserRequest struct {
	// ID int `json:`
	FullName string `json:"fullname" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Address  string `json:"address" gorm:"type: text" validate:"required"`
}
