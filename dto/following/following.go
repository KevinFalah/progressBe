package followingdto

type CreateFollowing struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type UpdateCollection struct {
	ID int `json:"id"`
}

type CollectionResponse struct {
	ID int `json:"id"`
}
