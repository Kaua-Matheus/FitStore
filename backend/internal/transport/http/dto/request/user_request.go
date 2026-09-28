package request

type RegisterRequest struct {
	UserName     string `json:"user_name" gorm:"not null"`
	UserLogin    string `json:"login" gorm:"not null"`
	UserPassword string `json:"password" gorm:"not null"`
}

type LoginRequest struct {
	UserLogin    string `json:"login" gorm:"not null"`
	UserPassword string `json:"password" gorm:"not null"`
}
