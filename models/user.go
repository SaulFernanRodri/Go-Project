package models

type User struct {
	ID           uint64 `gorm:"primaryKey;autoIncriment"`
	Name         string `gorm:"size:100;not null"`
	CSV          string `gorm:"size:1000"`
	AuthUsername string `gorm:"not null"`
}

type UserRequest struct {
	Name         string    `json:"name"`
	CSV          string    `json:"csv"`
	AuthUsername string    `json:"auth_username"`
	Milsymbol    Milsymbol `json:"mil_symbol"`
}
