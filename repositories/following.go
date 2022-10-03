package repositories

import (
	"waysgallery/models"

	"gorm.io/gorm"
)

type FollowingRepository interface {
	FindFollowings() ([]models.Following, error)
	GetFollowing(ID int) (models.Following, error)
	CreateFollowing(Following models.Following) (models.Following, error)
	DeleteFollowing(Following models.Following) (models.Following, error)
	GetFollow(ID int) (models.User, error)
}

func RepositoryFollowing(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFollowings() ([]models.Following, error) {
	var followings []models.Following
	err := r.db.Preload("User").Find(&followings).Error

	return followings, err
}

func (r *repository) GetFollowing(ID int) (models.Following, error) {
	var following models.Following
	err := r.db.Preload("User").First(&following, ID).Error

	return following, err
}

func (r *repository) CreateFollowing(following models.Following) (models.Following, error) {
	err := r.db.Preload("Following").Create(&following).Error

	return following, err
}

func (r *repository) UpdateFollowing(following []models.Following) ([]models.Following, error) {
	err := r.db.Save(&following).Error

	return following, err
}

func (r *repository) Updateefollowing(following models.Following) (models.Following, error) {
	err := r.db.Save(&following).Error

	return following, err
}

func (r *repository) DeleteFollowing(following models.Following) (models.Following, error) {
	err := r.db.Delete(&following).Error

	return following, err
}

func (r *repository) FindFollowingID([]int) ([]models.Following, error) {
	var following []models.Following
	err := r.db.Find(&following).Error

	return following, err
}

func (r *repository) GetFollow(ID int) (models.User, error) {
	var user models.User
	err := r.db.First(&user, ID).Error

	return user, err
}

// package repositories

// import (
// 	"waysgallery/models"

// 	"gorm.io/gorm"
// )

// type FollowingRepository interface {
// 	FindFollowings() ([]models.Following, error)
// 	GetFollowing(ID int) (models.Following, error)
// 	CreateFollowing(Following models.Following) (models.Following, error)
// 	UpdateFollowing(Following models.Following) (models.Following, error)
// 	DeleteFollowing(Following models.Following) (models.Following, error)
// 	GetFollowingPosts(ID int) ([]models.Post, error)
// }

// func RepositoryFollowing(db *gorm.DB) *repository {
// 	return &repository{db}
// }

// func (r *repository) FindFollowings() ([]models.Following, error) {
// 	var followings []models.Following
// 	err := r.db.Preload("Posts").Find(&followings).Error

// 	return followings, err
// }

// func (r *repository) GetFollowing(ID int) (models.Following, error) {
// 	var following models.Following
// 	err := r.db.Preload("Posts").First(&following, ID).Error

// 	return following, err
// }

// func (r *repository) CreateFollowing(following models.Following) (models.Following, error) {
// 	err := r.db.Create(&following).Error

// 	return following, err
// }

// func (r *repository) UpdateFollowing(following models.Following) (models.Following, error) {
// 	err := r.db.Save(&following).Error

// 	return following, err
// }

// func (r *repository) DeleteFollowing(following models.Following) (models.Following, error) {
// 	err := r.db.Delete(&following).Error

// 	return following, err
// }

// // func (r *repository) GetFollowingPosts(ID int) ([]models.Post, error) {
// // 	var posts []models.Post

// // 	err := r.db.Find(&posts, "following_id=?", ID).Error

// // 	return posts, err
// // }
