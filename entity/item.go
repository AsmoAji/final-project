package entity

import "time"

type Item struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"not null;unique"`
	Category     string    `gorm:"not null"`
	SerialNumber string    `gorm:"not null;unique"`
	PurchaseDate time.Time `gorm:"not null"`
	Status       string    `gorm:"default:'Aktif'"`
	Stock        int       `gorm:"default:0"`
	CreatedBy    uint
}
