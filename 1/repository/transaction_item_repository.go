package repository

import (
	"context"
	"myapp/model"

	"gorm.io/gorm"
)

type TransactionItemRepository interface {
	// insert
	Insert(ctx context.Context, transactionItem *model.TransactionItem) error
	InsertMany(ctx context.Context, transactionItems []model.TransactionItem) error

	// read
	Count(ctx context.Context) (int, error)
	Fetch(ctx context.Context) ([]model.TransactionItem, error)
	Get(ctx context.Context, id string) (*model.TransactionItem, error)

	// delete
	Delete(ctx context.Context, transactionItem *model.TransactionItem) error
}

type transactionItemRepository struct {
	gormDB *gorm.DB
}

func NewTransactionItemRepository(
	gormDB *gorm.DB,
) TransactionItemRepository {
	return &transactionItemRepository{
		gormDB: gormDB,
	}
}

func (r *transactionItemRepository) Insert(ctx context.Context, transactionItem *model.TransactionItem) error {
	result := r.gormDB.WithContext(ctx).Create(transactionItem)
	return result.Error
}

func (r *transactionItemRepository) InsertMany(ctx context.Context, transactionItems []model.TransactionItem) error {
	if len(transactionItems) == 0 {
		return nil
	}

	result := r.gormDB.WithContext(ctx).Create(transactionItems)
	return result.Error
}

func (r *transactionItemRepository) Count(ctx context.Context) (int, error) {
	var count int64
	result := r.gormDB.WithContext(ctx).Model(&model.TransactionItem{}).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *transactionItemRepository) Fetch(ctx context.Context) ([]model.TransactionItem, error) {
	var transactionItems []model.TransactionItem
	result := r.gormDB.WithContext(ctx).Find(&transactionItems)

	if result.Error != nil {
		return nil, result.Error
	}

	return transactionItems, nil
}

func (r *transactionItemRepository) Get(ctx context.Context, id string) (*model.TransactionItem, error) {
	var transactionItem model.TransactionItem
	result := r.gormDB.WithContext(ctx).First(&transactionItem, "id = ?", id)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &transactionItem, nil
}

func (r *transactionItemRepository) Delete(ctx context.Context, transactionItem *model.TransactionItem) error {
	result := r.gormDB.WithContext(ctx).Delete(transactionItem)
	return result.Error
}
