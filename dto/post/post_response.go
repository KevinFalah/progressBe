package postdto

type PostResponse struct {
	ID          int    `json:"id" gorm:"primary_key:auto_increment"`
	Title       string `json:"title" gorm:"type: varchar(255)"`
	Description string `json:"description" gorm:"type: varchar(255)"`
	Photo       string `json:"photo"`
	// UserID      int    `json:"user_id"`
	// User UsersResponse `json:"user"`
}
