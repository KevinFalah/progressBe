package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	FindPosts() ([]models.Post, error)
	GetPost(ID int) (models.Post, error)
	CreatePost(Post models.Post) (models.Post, error)
	UpdatePost(Post models.Post) (models.Post, error)
	DeletePost(Post models.Post) (models.Post, error)
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindPosts() ([]models.Post, error) {
	var posts []models.Post
	err := r.db.Preload("User").Find(&posts).Error // add this code

	return posts, err
}

func (r *repository) GetPost(ID int) (models.Post, error) {
	var Post models.Post
	//
	err := r.db.Preload("User").First(&Post, ID).Error // add this code

	return Post, err
}

func (r *repository) CreatePost(Post models.Post) (models.Post, error) {
	err := r.db.Create(&Post).Error

	return Post, err
}

func (r *repository) UpdatePost(Post models.Post) (models.Post, error) {
	err := r.db.Save(&Post).Error

	return Post, err
}

func (r *repository) DeletePost(Post models.Post) (models.Post, error) {
	err := r.db.Delete(&Post).Error

	return Post, err
}
