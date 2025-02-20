package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type ProductUseCase interface {
	Fetch(ctx context.Context, request dto_request.ProductFetchRequest) ([]model.Product, int, *dto_response.ErrorResponse, error)
	Get(ctx context.Context, request dto_request.ProductGetRequest) (*model.Product, *dto_response.ErrorResponse, error)
}

type productUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewProductUseCase(
	repositoryManager repository.RepositoryManager,
) ProductUseCase {
	return &productUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *productUseCase) Fetch(ctx context.Context, request dto_request.ProductFetchRequest) ([]model.Product, int, *dto_response.ErrorResponse, error) {
	queryOption := model.ProductQueryOption{
		QueryOption: model.QueryOption{
			Page:  request.Page,
			Limit: request.Limit,
		},
		Search: request.Search,
	}
	products, err := u.repositoryManager.ProductRepository().Fetch(ctx, queryOption)
	if err != nil {
		return nil, 0, nil, util.NewInternalServerError(err.Error())
	}

	total, err := u.repositoryManager.ProductRepository().Count(ctx, queryOption)
	if err != nil {
		return nil, 0, nil, util.NewInternalServerError(err.Error())
	}

	return products, total, nil, nil
}

func (u *productUseCase) Get(ctx context.Context, request dto_request.ProductGetRequest) (*model.Product, *dto_response.ErrorResponse, error) {
	product, err := u.repositoryManager.ProductRepository().Get(ctx, request.Id)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}
	if product == nil {
		return nil, dto_response.NewNotFoundErrorResponseP("Product not found"), nil
	}

	return product, nil, nil
}
