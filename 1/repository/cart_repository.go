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
	Count(ctx context.Context, options ...model.CartQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.CartQueryOption) ([]model.Cart, error)
	Get(ctx context.Context, id string) (*model.Cart, error)
	GetByUserIdAndProductId(ctx context.Context, userId string, productId string) (*model.Cart, error)

	// update
	Update(ctx context.Context, cart *model.Cart) error

	// delete
	Delete(ctx context.Context, cart *model.Cart) error
	DeleteMany(ctx context.Context, carts []model.Cart) error
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

func (r *cartRepository) Count(ctx context.Context, options ...model.CartQueryOption) (int, error) {
	option := model.CartQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var count int64
	stmt := r.gormDB.WithContext(ctx).Model(&model.Cart{})

	if option.UserId != nil {
		stmt = stmt.Where("user_id = ?", *option.UserId)
	}

	result := stmt.Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *cartRepository) Fetch(ctx context.Context, options ...model.CartQueryOption) ([]model.Cart, error) {
	option := model.CartQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var carts []model.Cart
	stmt := r.gormDB.WithContext(ctx)

	if option.UserId != nil {
		stmt = stmt.Where("user_id = ?", *option.UserId)
	}

	if option.LoadProduct {
		stmt = stmt.Preload("Product")
	}

	result := stmt.Find(&carts)

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

func (r *cartRepository) GetByUserIdAndProductId(ctx context.Context, userId string, productId string) (*model.Cart, error) {
	var cart model.Cart
	result := r.gormDB.WithContext(ctx).First(&cart, "user_id = ? AND product_id = ?", userId, productId)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &cart, nil
}

func (r *cartRepository) Update(ctx context.Context, cart *model.Cart) error {
	result := r.gormDB.WithContext(ctx).Updates(cart)
	return result.Error
}

func (r *cartRepository) Delete(ctx context.Context, cart *model.Cart) error {
	result := r.gormDB.WithContext(ctx).Delete(cart)
	return result.Error
}

func (r *cartRepository) DeleteMany(ctx context.Context, carts []model.Cart) error {
	result := r.gormDB.WithContext(ctx).Delete(carts)
	return result.Error
}
