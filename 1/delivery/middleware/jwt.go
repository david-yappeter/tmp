package middleware

import (
	"myapp/model"
	"myapp/use_case"

	"github.com/gin-gonic/gin"
)

func JWTHandler(router gin.IRouter, authUseCase use_case.AuthUseCase) {
	router.Use(func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			// add a checking get token from cookie 'access_token' first
			cookie, err := ctx.Request.Cookie("access_token")
			if err == nil {
				token = cookie.Value
			}
		}

		if token == "" {
			ctx.Next()
			return
		}

		user, clientError, err := authUseCase.Parse(ctx.Request.Context(), token)
		if err != nil {
			ctx.Next()
			return
		}
		if clientError != nil {
			ctx.Next()
			return
		}

		ctx.Request = ctx.Request.WithContext(model.SetUserCtx(ctx.Request.Context(), user))

		ctx.Next()
	})
}
