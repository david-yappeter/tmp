package model

import (
	"time"
)

const CartTableName = "carts"

type Cart struct {
	Id        string    `gorm:"primaryKey"`
	UserId    string    `gorm:"column:user_id"`
	ProductId string    `gorm:"column:product_id"`
	Qty       int       `gorm:"column:qty"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	Product *Product `gorm:"foreignKey:product_id"`
}

func (u *Cart) TableName() string {
	return CartTableName
}

type CartQueryOption struct {
	UserId *string

	LoadProduct bool
}
