package userdto

// type UserResponse struct {
// 	// ID       int    `json:"id"`
// 	FullName string `json:"fullname" form:"name" validate:"required"`
// 	Email    string `json:"email" form:"email" validate:"required"`
// 	Password string `json:"-" form:"name" validate:"required"`
// 	Phone    int    `json:"phone" gorm:"type: int(100)"`
// 	Address  string `json:"address" gorm:"type: text"`
// }

type DeleteUserResponse struct {
	ID int `json:"id"`
	// FullName string `json:"fullname" form:"name" validate:"required"`
	// Email    string `json:"email" form:"email" validate:"required"`
	// Password string `json:"-" form:"name" validate:"required"`
	// Phone    int    `json:"phone" gorm:"type: int(100)"`
	// Address  string `json:"address" gorm:"type: text"`
}
