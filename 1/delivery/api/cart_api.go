package api

import (
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"myapp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartApi struct {
	cartUseCase use_case.CartUseCase
}

func (a *CartApi) Create() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.CartCreateRequest
			if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
				return
			}

			cart, clientError, err := a.cartUseCase.Create(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.DataResponse{
				"cart": dto_response.NewCartResponse(*cart),
			})
		},
	)
}

func (a *CartApi) Fetch() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.CartFetchRequest
			if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
				return
			}

			carts, clientError, err := a.cartUseCase.Fetch(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.DataResponse{
				"carts": util.ConvertArray(carts, dto_response.NewCartResponse),
			})
		},
	)
}

func (a *CartApi) Get() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.CartGetRequest
			request.Id = ctx.Param("id")
			if !util.IsUuid(request.Id) {
				ctx.JSON(http.StatusBadRequest, dto_response.NewBadRequestErrorResponse("not a valid uuid"))
				return
			}

			cart, clientError, err := a.cartUseCase.Get(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.NewCartResponse(*cart))
		},
	)
}

func (a *CartApi) Update() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.CartUpdateRequest
			if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
				return
			}

			request.Id = ctx.Param("id")
			if !util.IsUuid(request.Id) {
				ctx.JSON(http.StatusBadRequest, dto_response.NewBadRequestErrorResponse("not a valid uuid"))
				return
			}

			cart, clientError, err := a.cartUseCase.Update(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.DataResponse{
				"cart": dto_response.NewCartResponse(*cart),
			})
		},
	)
}

func (a *CartApi) Delete() gin.HandlerFunc {
	return Authorize(
		func(ctx *gin.Context) {
			var request dto_request.CartDeleteRequest
			request.Id = ctx.Param("id")
			if !util.IsUuid(request.Id) {
				ctx.JSON(http.StatusBadRequest, dto_response.NewBadRequestErrorResponse("not a valid uuid"))
				return
			}

			clientError, err := a.cartUseCase.Delete(ctx.Request.Context(), request)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			if clientError != nil {
				ctx.JSON(clientError.Code, clientError)
				return
			}

			ctx.JSON(http.StatusOK, dto_response.DataResponse{
				"message": "OK",
			})
		},
	)
}

func RegisterCartApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CartApi{
		cartUseCase: useCaseManager.CartUseCase(),
	}
	routerGroup := router.Group("/carts")

	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}
