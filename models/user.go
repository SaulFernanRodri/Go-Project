package models

type User struct {
	ID          uint64    `gorm:"primaryKey;auto_incriment" json:"id"`
	Name        string    `gorm:"size:100" json:"name" binding:"required"`
	Age         int       `json:"age" binding:"gte=1,lte=100"`
	Email       string    `gorm:"unique" json:"email" binding:"required,email"`
	Password    string    `gorm:"size:100" json:"password" binding:"required,min=6"`
	Milsymbol   Milsymbol `gorm:"foreignKey:MilsymbolID" json:"milsymbol" binding:"required"`
	MilsymbolID uint64    `json:"milsymbol_id"`
	CSV         string    `gorm:"size:100" json:"csv"`
}
