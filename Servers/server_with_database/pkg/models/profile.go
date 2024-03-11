package models

import "time"

type Profile struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true" json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt time.Time `gorm:"autoUpdateTime" json:"deletedAt"`
}
