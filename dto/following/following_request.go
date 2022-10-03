package followingdto

type FollowingRequest struct {
	// UserID      int `json:"user_id" form:"user_id" gorm:"type: varchar(255)" validate:"required"`
	FollowingID int `json:"following_id" form:"following_id" gorm:"type: int" validate:"required"`
}
