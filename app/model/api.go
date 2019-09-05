package model

type Api struct {
	BaseModel
	Method  string
	Url     string
	GroupID uint
	Group   Group `gorm:"foreignkey:GroupID; association_foreignkey:ID"`
}

func (a Api) TableName() string {
	return "api"
}
