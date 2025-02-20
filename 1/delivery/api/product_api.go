package api

import (
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"myapp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductApi struct {
	productUseCase use_case.ProductUseCase
}

func (a *ProductApi) Fetch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto_request.ProductFetchRequest
		if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
			return
		}

		products, total, clientError, err := a.productUseCase.Fetch(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if clientError != nil {
			ctx.JSON(clientError.Code, clientError)
			return
		}

		ctx.JSON(http.StatusOK, dto_response.PaginationResponse{
			Page:  request.Page,
			Limit: request.Limit,
			Total: total,
			Nodes: util.ConvertArray(products, dto_response.NewProductResponse),
		})
	}
}

func (a *ProductApi) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto_request.ProductGetRequest
		request.Id = ctx.Param("id")

		if !util.IsUuid(request.Id) {
			ctx.JSON(http.StatusBadRequest, dto_response.NewBadRequestErrorResponse("not a valid uuid"))
			return
		}

		product, clientError, err := a.productUseCase.Get(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if clientError != nil {
			ctx.JSON(clientError.Code, clientError)
			return
		}

		ctx.JSON(http.StatusOK, dto_response.NewProductResponse(*product))
	}
}

func RegisterProductApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductApi{
		productUseCase: useCaseManager.ProductUseCase(),
	}
	routerGroup := router.Group("/products")

	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
}
