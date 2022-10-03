package models

import "time"

type Hired struct {
	ID             int                `json:"id" gorm:"primary_key:auto_increment"`
	Title          string             `json:"title" gorm:"type: varchar(255)" form:"title"`
	DescriptionJob string             `json:"descriptionjob" gorm:"type: varchar(255)" form:"descriptionjob"`
	StartProject   string             `json:"startproject" gorm:"type: varchar(255)" form:"startproject"`
	EndProject     string             `json:"endproject" gorm:"type: varchar(255)" form:"endproject"`
	Price          string             `json:"price" gorm:"type: varchar(255)" form:"price"`
	UserID         int                `json:"user_id" form:"user_id"`
	User           UsersHiredResponse `json:"user"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderToID      int                `json:"orderto_id" form:"orderto_id" gorm:"type:int"`
	OrderTo        OrderToHired       `json:"orderto" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt      time.Time          `json:"-"`
	UpdatedAt      time.Time          `json:"-"`
}


type HiredResponse struct {
	ID             int                `json:"id"`
	Title          string             `json:"title"`
	DescriptionJob string             `json:"descriptionjob"`
	StartProject   string             `json:"startproject"`
	EndProject     string             `json:"endproject"`
	Price          string             `json:"price"`
	UserID         int                `json:"-"`
	User           UsersHiredResponse `json:"user"`
	OrderToID      int                `json:"-" form:"orderto_id" gorm:"-"`
	OrderTo        OrderToHired       `json:"orderto"`
	CreatedAt      time.Time          `json:"-"`
	UpdatedAt      time.Time          `json:"-"`
}

func (HiredResponse) TableName() string {
	return "hireds"
}
