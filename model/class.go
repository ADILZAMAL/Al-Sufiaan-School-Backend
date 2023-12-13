package model

import (
	"time"
)

type Class struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique; not null" json:"name"`
	SchoolId  uint   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"school_id"`
	School    School `gorm:"foreignKey:SchoolId; references:Id"`
	Session   string `json:"session"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
