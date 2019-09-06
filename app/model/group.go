package model

type Group struct {
	Base
	Name  string `gorm:"UNIQUE; NOT NULL" json:"name"`
	Users []User `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"-"`
	Roles []Role `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"-"`
	Apis  []Api  `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"-"`
}

func (g Group) TableName() string {
	return "group"
}
