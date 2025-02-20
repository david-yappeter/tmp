package dto_response

import (
	"myapp/model"
	"time"
)

type ProductResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
} // @name ProductResponse

func NewProductResponse(product model.Product) ProductResponse {
	return ProductResponse{
		Id:        product.Id,
		Name:      product.Name,
		Price:     product.Price,
		CreatedAt: product.CreatedAt,
		UpdatedAt: product.UpdatedAt,
	}
}
