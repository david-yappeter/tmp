package dto_response

import (
	"myapp/model"
	"time"
)

type TransactionResponse struct {
	Id         string    `json:"id"`
	UserId     string    `json:"user_id"`
	TotalPrice float64   `json:"total_price"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	TransactionItems []TransactionItemResponse `json:"transaction_items"`
} // @name TransactionResponse

func NewTransactionResponse(transaction model.Transaction) TransactionResponse {
	r := TransactionResponse{
		Id:         transaction.Id,
		UserId:     transaction.UserId,
		TotalPrice: transaction.TotalPrice,
		CreatedAt:  transaction.CreatedAt,
		UpdatedAt:  transaction.UpdatedAt,
	}

	for _, transactionItem := range transaction.TransactionItems {
		r.TransactionItems = append(r.TransactionItems, NewTransactionItemResponse(transactionItem))
	}

	return r
}
