package model

type Group struct {
	BaseModel
	Name  string `gorm:"UNIQUE; NOT NULL" json:"name"`
	Users []User `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"users"`
}

func (g Group) TableName() string {
	return "group"
}
