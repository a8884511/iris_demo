package model

import "time"

type BaseModel struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedBy uint       `json:"created_by"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedBy uint       `json:"updated_by"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}
