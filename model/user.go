package model

import "time"

// User struct
type User struct {
	Id          uint   `gorm:"primary key"`
	Name        string `gorm:"not null" json:"name"`
	Email       string `gorm:"unique; not null" json:"email"`
	Password    string `json:"password"`
	Designation string `json:"designation"`
	SchoolId    uint   `gorm:"not null" json:"school_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
