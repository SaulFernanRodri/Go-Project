package models

type User struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"size:100" json:"password"`
	Milsymbol Milsymbol `gorm:"foreignKey:ID" json:"milsymbol"`
	CSV       string    `gorm:"size:100" json:"csv"`
}
