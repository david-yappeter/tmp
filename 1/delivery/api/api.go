package api

import (
	"context"
	"encoding/json"
	"io"
	"myapp/delivery/dto_response"
	"myapp/global"
	"myapp/manager"
	"myapp/use_case"
	"net/http"
	"time"

	bindingInternal "myapp/internal/gin/binding"
	"myapp/internal/gin/validator"
	_ "myapp/internal/gin/validator"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type apiContext struct {
	ginCtx *gin.Context
}

func newApiContext(ctx *gin.Context) apiContext {
	return apiContext{
		ginCtx: ctx,
	}
}

func (a *apiContext) context() context.Context {
	return a.ginCtx.Request.Context()
}

func NewRouter(container *manager.Container) *gin.Engine {
	allowedHeaders := []string{
		"Accept",
		"Accept-Encoding",
		"Authorization",
		"Cache-Control",
		"Content-Type",
		"Content-Length",
		"Origin",
		"X-CSRF-Token",
		"X-Requested-With",
	}

	if global.IsProduction() || global.IsTesting() {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(
		cors.New(
			cors.Config{
				AllowOrigins: global.GetConfig().CorsAllowedOrigins,
				AllowMethods: []string{
					http.MethodGet,
					http.MethodPost,
					http.MethodPut,
					http.MethodDelete,
					http.MethodPatch,
					http.MethodHead,
				},
				AllowHeaders: allowedHeaders,
				ExposeHeaders: []string{
					"Content-Type",
					"Content-Length",
					"Content-Disposition",
				},
				AllowCredentials: true,
				MaxAge:           2 * time.Hour,
			},
		),
	)

	registerMiddlewares(router, container)
	registerRoutes(router, container.UseCaseManager())

	return router
}

func registerMiddlewares(router gin.IRouter, container *manager.Container) {
	// useCaseManager := container.UseCaseManager()

	// middleware.JWTHandler(router, useCaseManager.AuthUseCase())
}

func registerRoutes(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	RegisterAuthApi(router, useCaseManager)
}

func translateBindErr(err error) dto_response.ErrorResponse {
	var r dto_response.ErrorResponse

	switch v := err.(type) {
	case validator.StructValidationErrors:
		errs := []dto_response.Error{}
		translations := v.Translate(validator.DefaultTranslator)
		for k, translation := range translations {
			errs = append(errs, dto_response.Error{
				Domain:  k,
				Message: translation,
			})
		}

		r = dto_response.NewBadRequestErrorResponse("Invalid request payload")
		r.Errors = errs

	case *json.UnmarshalTypeError, *json.InvalidUnmarshalError:
		r = dto_response.NewBadRequestErrorResponse("Invalid request payload, unmarshal error")

	default:
		switch v {
		case bindingInternal.ErrConvertMapStringSlice, bindingInternal.ErrConvertToMapString,
			bindingInternal.ErrMultiFileHeader, bindingInternal.ErrMultiFileHeaderLenInvalid,
			bindingInternal.ErrIgnoredBinding:
			r = dto_response.NewBadRequestErrorResponse("Invalid request payload")

		case io.EOF:
			r = dto_response.NewBadRequestErrorResponse("Empty Body")
		default:
			panic(err)
		}
	}

	return r
}
