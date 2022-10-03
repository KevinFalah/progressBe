package models

import "time"

type Following struct {
	ID          int                `json:"id" gorm:"primary_key:auto_increment"`
	UserID      int                `json:"user_id" form:"user_id"`
	User        UsersHiredResponse `json:"user"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FollowingID int                `json:"following_id" form:"following_id" gorm:"type:int"`
	Following   OrderToHired       `json:"following" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time          `json:"-"`
	UpdatedAt   time.Time          `json:"-"`
}

type FollowingResponse struct {
	ID          int                `json:"id"`
	UserID      int                `json:"user_id"`
	User        UsersHiredResponse `json:"user"`
	FollowingID int                `json:"-" form:"orderto_id" gorm:"-"`
	Following   OrderToHired       `json:"following"`
}

func (FollowingResponse) TableName() string {
	return "followings"
}
