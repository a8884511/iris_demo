package model

type Api struct {
	Base
	Method  string  `json:"method"`
	Url     string  `json:"url"`
	GroupID uint    `json:"group_id"`
	Group   Group   `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"-"`
	Roles   []*Role `gorm:"many2many:role_apis" json:"-"`
}

func (a Api) TableName() string {
	return "api"
}
