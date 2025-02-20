package model

import "time"

const TransactionItemTableName = "transaction_items"

type TransactionItem struct {
	Id           string    `gorm:"primaryKey"`
	ProductId    string    `gorm:"column:product_id"`
	PricePerUnit float64   `gorm:"column:price_per_unit"`
	CreatedAt    time.Time `gorm:"column:created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at"`
}

func (u *TransactionItem) TableName() string {
	return TransactionItemTableName
}
