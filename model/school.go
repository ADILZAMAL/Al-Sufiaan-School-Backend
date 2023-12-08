package model

import "time"

// School struct
type School struct {
	Id        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not	null" json:"name"`
	Address   string `json:"address"`
	Mobile    string `json:"mobile"`
	UdiceNo   string `gorm:"unique" json:"udice_no"`
	Active    bool
	Email     string `json:"email"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
