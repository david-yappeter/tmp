package dto_request

type CartCreateRequest struct {
	ProductId string `json:"product_id" binding:"required,uuid"`
}

type CartFetchRequest struct {
}

type CartGetRequest struct {
	Id string `json:"-"`
}

type CartUpdateRequest struct {
	Qty int `json:"qty" binding:"gt=0"`

	Id string `json:"-"`
}

type CartDeleteRequest struct {
	Id string `json:"-"`
}
