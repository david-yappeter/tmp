package dto_response

import "net/http"

type Response struct {
	Data interface{} `json:"data"`
} // @name Response

type SuccessResponse struct {
	Message string `json:"message" example:"OK"`
} // @name SuccessResponse

type Error struct {
	Domain  string `json:"domain"`
	Message string `json:"message"`
} // @name Error

type ErrorResponse struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Errors  []Error `json:"errors"`
} // @name ErrorResponse

type DataResponse map[string]interface{} // @name DataResponse

type PaginationResponse struct {
	Total int         `json:"total" example:"24"`
	Page  *int        `json:"page" example:"1"`
	Limit *int        `json:"limit" example:"10"`
	Nodes interface{} `json:"nodes"`
} // @name PaginationResponse

// func NewPaginationResponse(nodes interface{}, total int, pageP *int, limitP *int) PaginationResponse {
// 	page := constant.PaginationDefaultPage
// 	if pageP != nil {
// 		page = *pageP
// 	}

// 	limit := constant.PaginationDefaultLimit
// 	if limitP != nil {
// 		limit = *limitP
// 	}

// 	return PaginationResponse{
// 		Nodes: nodes,
// 		Total: total,
// 		Page:  util.Pointer(page),
// 		Limit: util.Pointer(limit),
// 	}
// }

// func NewPaginationResponseP(nodes interface{}, total int, pageP *int, limitP *int) *PaginationResponse {
// 	r := NewPaginationResponse(nodes, total, pageP, limitP)
// 	return &r
// }

func NewUnauthorizedErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusUnauthorized,
		Message: message,
		Errors:  []Error{},
	}
}

func NewUnauthorizedErrorResponseP(message string) *ErrorResponse {
	r := NewUnauthorizedErrorResponse(message)
	return &r
}

func NewBadRequestErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusBadRequest,
		Message: message,
		Errors:  []Error{},
	}
}

func NewBadRequestErrorResponseP(message string) *ErrorResponse {
	r := NewBadRequestErrorResponse(message)
	return &r
}

func NewForbiddenErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusForbidden,
		Message: message,
		Errors:  []Error{},
	}
}

func NewForbiddenErrorResponseP(message string) *ErrorResponse {
	r := NewForbiddenErrorResponse(message)
	return &r
}

func NewNotFoundErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusNotFound,
		Message: message,
		Errors:  []Error{},
	}
}

func NewNotFoundErrorResponseP(message string) *ErrorResponse {
	r := NewNotFoundErrorResponse(message)
	return &r
}

func NewConflictErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusConflict,
		Message: message,
		Errors:  []Error{},
	}
}

func NewConflictErrorResponseP(message string) *ErrorResponse {
	r := NewConflictErrorResponse(message)

	return &r
}

func NewInternalServerErrorResponse() ErrorResponse {
	return ErrorResponse{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error",
		Errors:  []Error{},
	}
}

func NewInternalServerErrorResponseP() *ErrorResponse {
	r := NewInternalServerErrorResponse()
	return &r
}
