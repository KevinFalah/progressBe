package models

import "time"

type Post struct {
	ID          int           `json:"id" gorm:"primary_key:auto_increment"`
	Title       string        `json:"title" gorm:"type: varchar(255)" form:"title"`
	Description string        `json:"description" gorm:"type: varchar(255)" form:"description"`
	Photo       string        `json:"photo" gorm:"type: varchar(255)" form:"photo"`
	UserID      int           `json:"user_id" form:"user_id"`
	User        UsersResponse `json:"user" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// User 		User	`json:"user"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type PostResponse struct {
	ID          int           `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Photo       string        `json:"photo"`
	UserID      int           `json:"-"`
	User        UsersResponse `json:"user"`
	// User       UsersResponse `json:"user"`
	// User 		User	`json:"user"`
}

type PostUserResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       string `json:"photo"`
	UserID      int    `json:"-"`
	// User 		User	`json:"user"`
}

func (PostResponse) TableName() string {
	return "posts"
}
func (PostUserResponse) TableName() string {
	return "posts"
}

// type UserPost struct {
// 	ID int `json:"id"`
// }

// func (UserPost) TableName() string {
// 	return "posts"
// }
