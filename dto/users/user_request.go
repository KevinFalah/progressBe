package usersdto

type CreateUserRequest struct {
	FullName string `json:"fullName" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Avatar   string `json:"avatar" form:"avatar" gorm:"type: varchar(255)"`
}

type UpdateUserRequest struct {
	FullName string `json:"fullName" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
