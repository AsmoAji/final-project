package entity

import "time"

type Item struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"not null;unique" json:"name"`
	Category     string    `json:"category"`
	SerialNumber string    `gorm:"not null;unique" json:"serial_number"`
	PurchaseDate time.Time `gorm:"not null" json:"purchase_date"`
	Status       string    `json:"status"`
	Stock        int       `json:"stock"`
	CreatedBy    uint      `json:"created_by"`
}
