package repository

import (
	"context"
	"myapp/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	// insert
	Insert(ctx context.Context, cart *model.Cart) error
	InsertMany(ctx context.Context, carts []model.Cart) error

	// read
	Count(ctx context.Context) (int, error)
	Fetch(ctx context.Context) ([]model.Cart, error)
	Get(ctx context.Context, id string) (*model.Cart, error)

	// delete
	Delete(ctx context.Context, cart *model.Cart) error
}

type cartRepository struct {
	gormDB *gorm.DB
}

func NewCartRepository(
	gormDB *gorm.DB,
) CartRepository {
	return &cartRepository{
		gormDB: gormDB,
	}
}

func (r *cartRepository) Insert(ctx context.Context, cart *model.Cart) error {
	result := r.gormDB.WithContext(ctx).Create(cart)
	return result.Error
}

func (r *cartRepository) InsertMany(ctx context.Context, carts []model.Cart) error {
	if len(carts) == 0 {
		return nil
	}

	result := r.gormDB.WithContext(ctx).Create(carts)
	return result.Error
}

func (r *cartRepository) Count(ctx context.Context) (int, error) {
	var count int64
	result := r.gormDB.WithContext(ctx).Model(&model.Cart{}).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *cartRepository) Fetch(ctx context.Context) ([]model.Cart, error) {
	var carts []model.Cart
	result := r.gormDB.WithContext(ctx).Find(&carts)

	if result.Error != nil {
		return nil, result.Error
	}

	return carts, nil
}

func (r *cartRepository) Get(ctx context.Context, id string) (*model.Cart, error) {
	var cart model.Cart
	result := r.gormDB.WithContext(ctx).First(&cart, "id = ?", id)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &cart, nil
}

func (r *cartRepository) Delete(ctx context.Context, cart *model.Cart) error {
	result := r.gormDB.WithContext(ctx).Delete(cart)
	return result.Error
}
