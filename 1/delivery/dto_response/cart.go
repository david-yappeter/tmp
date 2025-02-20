package dto_response

import (
	"myapp/model"
	"time"
)

type CartResponse struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	ProductId string    `json:"product_id"`
	Qty       int       `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCartResponse(cart model.Cart) CartResponse {
	r := CartResponse{
		Id:        cart.Id,
		UserId:    cart.UserId,
		ProductId: cart.ProductId,
		Qty:       cart.Qty,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}

	return r
}
