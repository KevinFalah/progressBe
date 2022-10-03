package models

import "time"

type User struct {
	ID         int                 `json:"id"`
	FullName   string              `json:"fullName" form:"name" validate:"required"`
	Email      string              `json:"email" form:"email" validate:"required"`
	Password   string              `json:"password" form:"password" validate:"required"`
	Avatar     string              `json:"avatar" form:"avatar" gorm:"type: varchar(255)"`
	Greeting   string              `json:"greeting" form:"greeting" gorm:"type: varchar(255)"`
	Posts      []PostUserResponse  `json:"posts"`
	Followings []FollowingResponse `json:"following"`
	Hireds     []HiredResponse     `json:"hired"`
	CreatedAt  time.Time           `json:"-"`
	UpdatedAt  time.Time           `json:"-"`
}

type UsersHiredResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type OrderToHired struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
}

type UsersResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Greeting string `json:"greeting"`
	// FullName string `json:"fullname"`
}

func (UsersHiredResponse) TableName() string {
	return "users"
}

func (OrderToHired) TableName() string {
	return "users"
}

func (UsersResponse) TableName() string {
	return "users"
}
