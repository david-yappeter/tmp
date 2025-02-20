package dto_response

import (
	"myapp/model"
	"time"
)

type TransactionItemResponse struct {
	Id            string    `json:"id"`
	TransactionId string    `json:"transaction_id"`
	ProductId     string    `json:"product_id"`
	PricePerUnit  float64   `json:"price_per_unit"`
	Qty           int       `json:"qty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
} // @name TransactionItemResponse

func NewTransactionItemResponse(transactionItem model.TransactionItem) TransactionItemResponse {
	r := TransactionItemResponse{
		Id:            transactionItem.Id,
		TransactionId: transactionItem.TransactionId,
		ProductId:     transactionItem.ProductId,
		PricePerUnit:  transactionItem.PricePerUnit,
		Qty:           transactionItem.Qty,
		CreatedAt:     transactionItem.CreatedAt,
		UpdatedAt:     transactionItem.UpdatedAt,
	}

	return r
}
