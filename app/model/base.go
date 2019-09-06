package model

import "time"

type Base struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedBy uint       `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedBy uint       `json:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
	Desc      string     `json:"desc"`
}

//这样封装的目的是为了以后可以更新其他扩展字段
func (b *Base) UpdateStatus(status map[string]interface{}) {
	b.Desc = status["desc"].(string)
}
