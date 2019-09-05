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

//角色名唯一，可以加上不同组的前缀
type Role struct {
	BaseModel
	Name    string
	GroupID uint
	Group   Group `gorm:"foreignkey:GroupID; association_foreignkey:ID"`
}

func (r Role) TableName() string {
	return "role"
}
