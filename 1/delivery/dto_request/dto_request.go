package dto_request

type PaginationRequest struct {
	Page  *int `json:"page" binding:"required_with=Limit,omitempty,gte=1"`
	Limit *int `json:"limit" binding:"required_with=Page,omitempty,gte=1,lte=100"`
}
