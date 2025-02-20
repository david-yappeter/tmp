package use_case

import (
	"myapp/infrastructure"
	jwtInternal "myapp/internal/jwt"
	"myapp/repository"
)

type UseCaseManager interface {
	AuthUseCase() AuthUseCase
	ProductUseCase() ProductUseCase
}

type useCaseManager struct {
	authUseCase    AuthUseCase
	productUseCase ProductUseCase
}

func (u *useCaseManager) AuthUseCase() AuthUseCase {
	return u.authUseCase
}

func (u *useCaseManager) ProductUseCase() ProductUseCase {
	return u.productUseCase
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
		productUseCase: NewProductUseCase(
			repositoryManager,
		),
	}
}
