package authdto

type LoginResponse struct {
	FullName string `gorm:"type: varchar(255)" json:"fullName"`
	Email    string `gorm:"type: varchar(255)" json:"email"`
	Password string `gorm:"type: varchar(255)" json:"password"`
	Token    string `gorm:"type: varchar(255)" json:"token"`
	UserID   int	`json:"user_id"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	Fullname  string `gorm:"type: varchar(255)" json:"fullName"`
	UserID   int	`json:"user_id"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Address   string `gorm:"type: varchar(255)" json:"address"`
	Token     string `gorm:"type: varchar(255)" json:"token"`
}