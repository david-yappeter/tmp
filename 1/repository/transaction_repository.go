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
	Count(ctx context.Context, options ...model.TransactionQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.TransactionQueryOption) ([]model.Transaction, error)
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

func (r *transactionRepository) Count(ctx context.Context, options ...model.TransactionQueryOption) (int, error) {
	option := model.TransactionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var count int64
	stmt := r.gormDB.WithContext(ctx).Model(&model.Transaction{})

	if option.UserId != nil {
		stmt = stmt.Where("user_id = ?", *option.UserId)
	}

	result := stmt.Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *transactionRepository) Fetch(ctx context.Context, options ...model.TransactionQueryOption) ([]model.Transaction, error) {
	option := model.TransactionQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var transactions []model.Transaction
	stmt := r.gormDB.WithContext(ctx)

	if option.UserId != nil {
		stmt = stmt.Where("user_id = ?", *option.UserId)
	}

	result := stmt.Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactions, nil
}

func (r *transactionRepository) Get(ctx context.Context, id string) (*model.Transaction, error) {
	var transaction model.Transaction
	result := r.gormDB.WithContext(ctx).Preload("TransactionItems").First(&transaction, "id = ?", id)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &transaction, nil
}

func (r *transactionRepository) Delete(ctx context.Context, transaction *model.Transaction) error {
	result := r.gormDB.WithContext(ctx).Delete(transaction)
	return result.Error
}
