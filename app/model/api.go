package model

type Api struct {
	Base
	Method  string `json:"method"`
	Url     string `json:"url"`
	GroupID uint   `json:"group_id"`
	Group   Group  `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"-"`
	RoleID  uint   `json:"role_id"`
	Role    Role   `gorm:"foreignkey:RoleID; association_foreignkey:ID" json:"-"`
}

func (a Api) TableName() string {
	return "api"
}
