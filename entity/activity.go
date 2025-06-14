package entity

import "time"

type Activity struct {
	ID     uint   `gorm:"primaryKey"`
	ItemID uint   `gorm:"not null"`
	UserID uint   `gorm:"not null"`
	Action string `gorm:"not null"` // contoh: 'pemeliharaan'
	Notes  string
	Date   time.Time `gorm:"not null"`
}
