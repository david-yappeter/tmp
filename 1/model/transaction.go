package model

import "time"

const TransactionTableName = "transactions"

type Transaction struct {
	Id               string            `gorm:"primaryKey"`
	UserId           string            `gorm:"column:user_id"`
	TotalPrice       float64           `gorm:"column:total_price"`
	CreatedAt        time.Time         `gorm:"column:created_at"`
	UpdatedAt        time.Time         `gorm:"column:updated_at"`
	TransactionItems []TransactionItem `gorm:"foreignKey:TransactionId"`
}

func (u *Transaction) TableName() string {
	return TransactionTableName
}

type TransactionQueryOption struct {
	UserId *string
}
