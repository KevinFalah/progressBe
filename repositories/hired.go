package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type HiredRepository interface {
	FindHireds() ([]models.Hired, error)
	GetHired(ID int) (models.Hired, error)
	CreateHired(hired models.Hired) (models.Hired, error)
	GetOrderTo(ID int) (models.User, error)
}

func RepositoryHired(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindHireds() ([]models.Hired, error) {
	var hireds []models.Hired
	err := r.db.Preload("User").Preload("OrderTo").Find(&hireds).Error

	return hireds, err
}

func (r *repository) GetHired(ID int) (models.Hired, error) {
	var hired models.Hired
	err := r.db.Preload("User").First(&hired, ID).Error

	return hired, err
}

func (r *repository) CreateHired(hired models.Hired) (models.Hired, error) {
	err := r.db.Create(&hired).Error

	return hired, err
}

func (r *repository) GetOrderTo(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}
