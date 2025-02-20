package api

import (
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"myapp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionApi struct {
	transactionUseCase use_case.TransactionUseCase
}

func (a *TransactionApi) Checkout() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.TransactionCheckoutRequest
			if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
				return
			}

			transaction, clientError, err := a.transactionUseCase.Checkout(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.DataResponse{
				"transaction": dto_response.NewTransactionResponse(*transaction),
			})
		},
	)
}

func (a *TransactionApi) Fetch() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.TransactionFetchRequest
			if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
				return
			}

			transactions, clientError, err := a.transactionUseCase.Fetch(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.DataResponse{
				"transactions": util.ConvertArray(transactions, dto_response.NewTransactionResponse),
			})
		},
	)
}

func (a *TransactionApi) Get() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.TransactionGetRequest
			request.Id = ctx.Param("id")
			if !util.IsUuid(request.Id) {
				ctx.JSON(http.StatusBadRequest, dto_response.NewBadRequestErrorResponse("not a valid uuid"))
				return
			}

			transaction, clientError, err := a.transactionUseCase.Get(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.NewTransactionResponse(*transaction))
		},
	)
}

func RegisterTransactionApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := TransactionApi{
		transactionUseCase: useCaseManager.TransactionUseCase(),
	}
	routerGroup := router.Group("/transactions")

	routerGroup.POST("/checkout", api.Checkout())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
}
