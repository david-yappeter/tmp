package api

import (
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	authUseCase use_case.AuthUseCase
}

func (a *AuthApi) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto_request.AuthLoginRequest
		if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
			return
		}

		token, clientError, err := a.authUseCase.Login(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if clientError != nil {
			ctx.JSON(clientError.Code, clientError)
			return
		}

		ctx.JSON(http.StatusOK, dto_response.NewAuthTokenResponse(*token))
	}
}

func (a *AuthApi) Register() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var request dto_request.AuthRegisterRequest
		if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": translateBindErr(err)})
			return
		}

		token, clientError, err := a.authUseCase.Register(ctx.Request.Context(), request)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		if clientError != nil {
			ctx.JSON(clientError.Code, clientError)
			return
		}

		ctx.JSON(http.StatusOK, dto_response.NewAuthTokenResponse(*token))
	}
}

func RegisterAuthApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AuthApi{
		authUseCase: useCaseManager.AuthUseCase(),
	}
	routerGroup := router.Group("/auth")

	routerGroup.POST("/login", api.Login())
	routerGroup.POST("/register", api.Register())
}
