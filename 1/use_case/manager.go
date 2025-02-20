package use_case

import (
	"myapp/infrastructure"
	jwtInternal "myapp/internal/jwt"
	"myapp/repository"
)

type UseCaseManager interface {
	AuthUseCase() AuthUseCase
	CartUseCase() CartUseCase
	ProductUseCase() ProductUseCase
	TransactionUseCase() TransactionUseCase
}

type useCaseManager struct {
	authUseCase        AuthUseCase
	cartUseCase        CartUseCase
	productUseCase     ProductUseCase
	transactionUseCase TransactionUseCase
}

func (u *useCaseManager) AuthUseCase() AuthUseCase {
	return u.authUseCase
}

func (u *useCaseManager) CartUseCase() CartUseCase {
	return u.cartUseCase
}

func (u *useCaseManager) ProductUseCase() ProductUseCase {
	return u.productUseCase
}

func (u *useCaseManager) TransactionUseCase() TransactionUseCase {
	return u.transactionUseCase
}

func NewUseCaseManager(
	infrastructureManager infrastructure.InfrastructureManager,
	repositoryManager repository.RepositoryManager,
	jwt jwtInternal.Jwt,
) UseCaseManager {
	return &useCaseManager{
		authUseCase: NewAuthUseCase(
			repositoryManager,
			jwt,
		),
		cartUseCase: NewCartUseCase(
			repositoryManager,
		),
		productUseCase: NewProductUseCase(
			repositoryManager,
		),
		transactionUseCase: NewTransactionUseCase(
			repositoryManager,
		),
	}
}
