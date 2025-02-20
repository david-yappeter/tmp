package model

import "time"

const ProductTableName = "products"

type Product struct {
	Id        string    `gorm:"primaryKey"`
	Name      string    `gorm:"column:name"`
	Price     float64   `gorm:"column:price"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (u *Product) TableName() string {
	return ProductTableName
}

type ProductQueryOption struct {
	QueryOption

	Search *string
}
