package models

import "time"

type Product struct {
	ProductID uint      `json:"productID" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	Quantity  int       `json:"quantity"`
	EntryDate time.Time `json:"entryDate"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt time.Time `gorm:"autoCreateTime" json:"deletedAt"`
	ProfileID uint      `json:"addedBy"`
	Profile   Profile   `gorm:"foreignkey:ProfileID;references:ID"`
}
