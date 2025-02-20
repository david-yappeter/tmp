package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type CartUseCase interface {
	Create(ctx context.Context, request dto_request.CartCreateRequest) (*model.Cart, *dto_response.ErrorResponse, error)

	Fetch(ctx context.Context, request dto_request.CartFetchRequest) ([]model.Cart, *dto_response.ErrorResponse, error)
	Get(ctx context.Context, request dto_request.CartGetRequest) (*model.Cart, *dto_response.ErrorResponse, error)

	Update(ctx context.Context, request dto_request.CartUpdateRequest) (*model.Cart, *dto_response.ErrorResponse, error)

	Delete(ctx context.Context, request dto_request.CartDeleteRequest) (*dto_response.ErrorResponse, error)
}

type cartUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewCartUseCase(
	repositoryManager repository.RepositoryManager,
) CartUseCase {
	return &cartUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *cartUseCase) Create(ctx context.Context, request dto_request.CartCreateRequest) (*model.Cart, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	product, err := u.repositoryManager.ProductRepository().Get(ctx, request.ProductId)
	if err != nil {
		return nil, nil, err
	}
	if product == nil {
		return nil, dto_response.NewNotFoundErrorResponseP("Product not found"), nil
	}

	var cart *model.Cart

	cart, err = u.repositoryManager.CartRepository().GetByUserIdAndProductId(ctx, authUser.Id, request.ProductId)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	if cart != nil {
		cart.Qty += 1

		err = u.repositoryManager.CartRepository().Update(ctx, cart)
		if err != nil {
			return nil, nil, util.NewInternalServerError(err.Error())
		}

	} else {
		cart = &model.Cart{
			Id:        util.NewUuid(),
			UserId:    authUser.Id,
			ProductId: request.ProductId,
			Qty:       1,
		}

		err = u.repositoryManager.CartRepository().Insert(ctx, cart)
		if err != nil {
			return nil, nil, util.NewInternalServerError(err.Error())
		}
	}

	return cart, nil, nil
}

func (u *cartUseCase) Fetch(ctx context.Context, request dto_request.CartFetchRequest) ([]model.Cart, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}
	queryOption := model.CartQueryOption{
		UserId: &authUser.Id,
	}

	carts, err := u.repositoryManager.CartRepository().Fetch(ctx, queryOption)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	return carts, nil, nil
}

func (u *cartUseCase) Get(ctx context.Context, request dto_request.CartGetRequest) (*model.Cart, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	cart, err := u.repositoryManager.CartRepository().Get(ctx, request.Id)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}
	if cart == nil || cart.UserId != authUser.Id {
		return nil, dto_response.NewNotFoundErrorResponseP("Cart not found"), nil
	}

	return cart, nil, nil
}

func (u *cartUseCase) Update(ctx context.Context, request dto_request.CartUpdateRequest) (*model.Cart, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	cart, err := u.repositoryManager.CartRepository().Get(ctx, request.Id)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}
	if cart == nil || cart.UserId != authUser.Id {
		return nil, dto_response.NewNotFoundErrorResponseP("Cart not found"), nil
	}

	cart.Qty = request.Qty

	err = u.repositoryManager.CartRepository().Update(ctx, cart)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	return cart, nil, nil
}

func (u *cartUseCase) Delete(ctx context.Context, request dto_request.CartDeleteRequest) (*dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, err
	}

	cart, err := u.repositoryManager.CartRepository().Get(ctx, request.Id)
	if err != nil {
		return nil, util.NewInternalServerError(err.Error())
	}
	if cart == nil || cart.UserId != authUser.Id {
		return dto_response.NewNotFoundErrorResponseP("Cart not found"), nil
	}

	err = u.repositoryManager.CartRepository().Delete(ctx, cart)
	if err != nil {
		return nil, util.NewInternalServerError(err.Error())
	}

	return nil, nil
}
