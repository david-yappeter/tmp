package dto_request

type ProductFetchRequest struct {
	PaginationRequest

	Search *string `json:"search"`
}

type ProductGetRequest struct {
	Id string `json:"-"`
}
