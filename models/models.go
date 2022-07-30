package models

import "gorm.io/gorm"

// Product model
type Product struct {
	gorm.Model     `swaggerignore:"true"`
	Name           string  `json:"name"  gorm:"column:name" example:"Great product"`
	Price          float32 `json:"price" gorm:"column:price" example:"20"`
	AvailableStock int32   `json:"stock" gorm:"column:available_stock" example:"10"`
}
