package hireddto

type HiredRequest struct {
	Title          string `json:"title" form:"title" gorm:"type: varchar(255)" validate:"required"`
	DescriptionJob string `json:"descriptionjob" gorm:"type:text" form:"descriptionjob" validate:"required"`
	StartProject   string `json:"startproject" gorm:"type:text" form:"startproject" validate:"required"`
	EndProject     string `json:"endproject" gorm:"type:text" form:"endproject" validate:"required"`
	Price          string `json:"price" gorm:"type:text" form:"price" validate:"required"`
	OrderToID      int    `json:"orderto_id" form:"orderto_id" gorm:"type: int" validate:"required"`
	
}
