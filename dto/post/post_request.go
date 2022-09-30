package postdto

// type Photo struct {
// 	Image string `json:"image" gorm:"type:varchar(255)"`
// }

type CreatePostRequest struct {
	Title       string `json:"title" gorm:"type:varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	Photo       string `json:"photo"`
	UserID      int    `json:"user_id"`
}

type PostRequest struct {
	Title       string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	Description string `json:"description" gorm:"type:text" form:"description" validate:"required"`
	// Photo       string `json:"photo"`
	// UserID      int    `json:"user_id"`
}

type UpdatePostRequest struct {
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type:text"`
	Photo       string `json:"photo"`
}

// type PostRequest struct {
// 	Title string `json:"title" gorm:"type:varchar(255)"`
// 	Description   string `json:"description" gorm:"type:text"`
// }
