package model

//foreignkey指定以哪个属性去关联
//A(.ID)->B

//association_foreignkey指定关联到模型的那个属性
//A->B(.ID)

//type Role struct {
//	gorm.Model
//	Name     string `gorm:"UNIQUE; NOT NULL"`
//  GroupID  uint
//	Group    Group  `gorm:"foreignkey:GroupID, association_foreignkey:ID"`
//}
//Role.GroupID == Group.ID

type Role struct {
	Base
	Name    string `gorm:"NOT NULL" json:"name"`
	GroupID uint   `json:"group_id"`
	Group   Group  `gorm:"foreignkey:GroupID; association_foreignkey:ID" json:"-"`
	Apis    []Api  `gorm:"foreignkey:RoleID; association_foreignkey:ID" json:"-"`
}

func (r Role) TableName() string {
	return "role"
}
