package entity

import "time"

type Activity struct {
	ID     uint      `gorm:"primaryKey" json:"id"`
	ItemID uint      `gorm:"not null" json:"item_id"`
	UserID uint      `gorm:"not null" json:"user_id"`
	Action string    `gorm:"not null" json:"action"` // contoh: 'pemeliharaan'
	Notes  string    `json:"notes"`
	Date   time.Time `gorm:"not null" json:"date"`
}
