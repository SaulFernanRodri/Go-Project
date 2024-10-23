package models

type InfoFields struct {
	UniqueDesignation string `gorm:"column:unique_designation" json:"uniqueDesignation"`
	HigherFormation   string `gorm:"column:higher_formation" json:"higherFormation"`
	StaffComments     string `gorm:"column:staff_comments" json:"staffComments"`
	Speed             string `gorm:"column:speed" json:"speed"`
}

type Milsymbol struct {
	ID         uint64     `gorm:"primaryKey;auto_incriment" json:"id"`
	SymbolCode string     `gorm:"column:symbol_code" json:"symbolcode" binding:"required"` // Campo requerido
	Size       int        `gorm:"column:size" json:"size" binding:"required"`              // Campo requerido
	Frame      bool       `gorm:"column:frame" json:"frame"`
	Fill       string     `gorm:"column:fill" json:"fill"`
	InfoFields InfoFields `gorm:"embedded" json:"infofields"`
	Quantity   int        `gorm:"column:quantity" json:"quantity"`
	Direction  int        `gorm:"column:direction" json:"direction"`
	Status     string     `gorm:"column:status" json:"status"`
}
