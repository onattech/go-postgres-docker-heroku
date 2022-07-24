package models

import "gorm.io/gorm"

// Product model
type Product struct {
	gorm.Model

	// ID             int64   `json:"id"    gorm:"column:id; primaryKey; <-:create"`
	Name           string  `json:"name"  gorm:"column:name"`
	Price          float32 `json:"price" gorm:"column:price"`
	AvailableStock int32   `json:"stock" gorm:"column:available_stock"`
}
