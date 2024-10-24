package models

type InfoFields struct {
	UniqueDesignation string `json:"uniqueDesignation"`
	HigherFormation   string `json:"higherFormation"`
	StaffComments     string `json:"staffComments"`
	Speed             string `json:"speed"`
}

type Milsymbol struct {
	SymbolCode string     `json:"symbolcode" binding:"required"` // Campo requerido
	Size       int        `json:"size" binding:"required"`       // Campo requerido
	Frame      bool       `json:"frame"`
	Fill       string     `json:"fill"`
	InfoFields InfoFields `json:"info_fields"`
	Quantity   int        `json:"quantity"`
	Direction  int        `json:"direction"`
	Status     string     `json:"status"`
}
