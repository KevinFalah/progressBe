package photodto

// type Photo struct {
// 	Image string `json:"image" gorm:"type:varchar(255)"`
// }

type PhotoResponse struct {
	Image string `json:"image"`
	PostID int 	`json:"post_id"`
}