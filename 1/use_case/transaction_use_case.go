package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type TransactionUseCase interface {
	Checkout(ctx context.Context, request dto_request.TransactionCheckoutRequest) (*model.Transaction, *dto_response.ErrorResponse, error)

	Fetch(ctx context.Context, request dto_request.TransactionFetchRequest) ([]model.Transaction, *dto_response.ErrorResponse, error)
	Get(ctx context.Context, request dto_request.TransactionGetRequest) (*model.Transaction, *dto_response.ErrorResponse, error)
}

type transactionUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewTransactionUseCase(
	repositoryManager repository.RepositoryManager,
) TransactionUseCase {
	return &transactionUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *transactionUseCase) Checkout(ctx context.Context, request dto_request.TransactionCheckoutRequest) (*model.Transaction, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	carts, err := u.repositoryManager.CartRepository().Fetch(ctx, model.CartQueryOption{
		UserId:      &authUser.Id,
		LoadProduct: true,
	})
	if err != nil {
		return nil, nil, err
	}
	if len(carts) == 0 {
		return nil, dto_response.NewNotFoundErrorResponseP("Carts Empty"), nil
	}

	transaction := model.Transaction{
		Id:         util.NewUuid(),
		UserId:     authUser.Id,
		TotalPrice: 0,
	}

	transactionItems := []model.TransactionItem{}
	for _, cart := range carts {
		subtotal := cart.Product.Price * float64(cart.Qty)
		transaction.TotalPrice += subtotal

		transactionItems = append(transactionItems, model.TransactionItem{
			Id:            util.NewUuid(),
			TransactionId: transaction.Id,
			ProductId:     cart.ProductId,
			PricePerUnit:  cart.Product.Price,
			Qty:           cart.Qty,
		})
	}

	err = u.repositoryManager.TransactionRepository().Insert(ctx, &transaction)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	err = u.repositoryManager.TransactionItemRepository().InsertMany(ctx, transactionItems)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	err = u.repositoryManager.CartRepository().DeleteMany(ctx, carts)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	transaction.TransactionItems = transactionItems

	return &transaction, nil, nil
}

func (u *transactionUseCase) Fetch(ctx context.Context, request dto_request.TransactionFetchRequest) ([]model.Transaction, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}
	queryOption := model.TransactionQueryOption{
		UserId: &authUser.Id,
	}

	transactions, err := u.repositoryManager.TransactionRepository().Fetch(ctx, queryOption)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}

	return transactions, nil, nil
}

func (u *transactionUseCase) Get(ctx context.Context, request dto_request.TransactionGetRequest) (*model.Transaction, *dto_response.ErrorResponse, error) {
	authUser, err := model.GetUserCtx(ctx)
	if err != nil {
		return nil, nil, err
	}

	transaction, err := u.repositoryManager.TransactionRepository().Get(ctx, request.Id)
	if err != nil {
		return nil, nil, util.NewInternalServerError(err.Error())
	}
	if transaction == nil || transaction.UserId != authUser.Id {
		return nil, dto_response.NewNotFoundErrorResponseP("Transaction not found"), nil
	}

	return transaction, nil, nil
}
