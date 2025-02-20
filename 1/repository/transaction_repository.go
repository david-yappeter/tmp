package repository

import (
	"context"
	"myapp/model"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	// insert
	Insert(ctx context.Context, transaction *model.Transaction) error
	InsertMany(ctx context.Context, transactions []model.Transaction) error

	// read
	Count(ctx context.Context) (int, error)
	Fetch(ctx context.Context) ([]model.Transaction, error)
	Get(ctx context.Context, id string) (*model.Transaction, error)

	// delete
	Delete(ctx context.Context, transaction *model.Transaction) error
}

type transactionRepository struct {
	gormDB *gorm.DB
}

func NewTransactionRepository(
	gormDB *gorm.DB,
) TransactionRepository {
	return &transactionRepository{
		gormDB: gormDB,
	}
}

func (r *transactionRepository) Insert(ctx context.Context, transaction *model.Transaction) error {
	result := r.gormDB.WithContext(ctx).Create(transaction)
	return result.Error
}

func (r *transactionRepository) InsertMany(ctx context.Context, transactions []model.Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	result := r.gormDB.WithContext(ctx).Create(transactions)
	return result.Error
}

func (r *transactionRepository) Count(ctx context.Context) (int, error) {
	var count int64
	result := r.gormDB.WithContext(ctx).Model(&model.Transaction{}).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *transactionRepository) Fetch(ctx context.Context) ([]model.Transaction, error) {
	var transactions []model.Transaction
	result := r.gormDB.WithContext(ctx).Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (r *transactionRepository) Get(ctx context.Context, id string) (*model.Transaction, error) {
	var transaction model.Transaction
	result := r.gormDB.WithContext(ctx).First(&transaction, "id = ?", id)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &transaction, nil
}

func (r *transactionRepository) Delete(ctx context.Context, transaction *model.Transaction) error {
	result := r.gormDB.WithContext(ctx).Delete(transaction)
	return result.Error
}
