package repository

import (
	"myapp/infrastructure"

	"gorm.io/gorm"
)

type RepositoryManager interface {
	CartRepository() CartRepository
	ProductRepository() ProductRepository
	TransactionRepository() TransactionRepository
	TransactionItemRepository() TransactionItemRepository
	UserRepository() UserRepository
}

type repositoryManager struct {
	gormDB *gorm.DB

	cartRepository            CartRepository
	productRepository         ProductRepository
	transactionRepository     TransactionRepository
	transactionItemRepository TransactionItemRepository
	userRepository            UserRepository
}

func (r *repositoryManager) CartRepository() CartRepository {
	return r.cartRepository
}

func (r *repositoryManager) ProductRepository() ProductRepository {
	return r.productRepository
}

func (r *repositoryManager) TransactionRepository() TransactionRepository {
	return r.transactionRepository
}

func (r *repositoryManager) TransactionItemRepository() TransactionItemRepository {
	return r.transactionItemRepository
}

func (r *repositoryManager) UserRepository() UserRepository {
	return r.userRepository
}

func NewRepositoryManager(infrastructureManager infrastructure.InfrastructureManager) RepositoryManager {
	gormDB := infrastructureManager.GetGormDB()

	return &repositoryManager{
		gormDB: gormDB,

		cartRepository: NewCartRepository(
			gormDB,
		),
		productRepository: NewProductRepository(
			gormDB,
		),
		transactionRepository: NewTransactionRepository(
			gormDB,
		),
		transactionItemRepository: NewTransactionItemRepository(
			gormDB,
		),
		userRepository: NewUserRepository(
			gormDB,
		),
	}
}
