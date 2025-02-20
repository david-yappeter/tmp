package repository

import (
	"context"
	"myapp/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	// insert
	Insert(ctx context.Context, product *model.Product) error
	InsertMany(ctx context.Context, products []model.Product) error

	// read
	Count(ctx context.Context, options ...model.ProductQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.ProductQueryOption) ([]model.Product, error)
	Get(ctx context.Context, id string) (*model.Product, error)
}

type productRepository struct {
	gormDB *gorm.DB
}

func NewProductRepository(
	gormDB *gorm.DB,
) ProductRepository {
	return &productRepository{
		gormDB: gormDB,
	}
}

func (r *productRepository) Insert(ctx context.Context, product *model.Product) error {
	result := r.gormDB.WithContext(ctx).Create(product)
	return result.Error
}

func (r *productRepository) InsertMany(ctx context.Context, products []model.Product) error {
	if len(products) == 0 {
		return nil
	}

	result := r.gormDB.WithContext(ctx).Create(products)
	return result.Error
}

func (r *productRepository) Count(ctx context.Context, options ...model.ProductQueryOption) (int, error) {
	option := model.ProductQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var count int64
	stmt := r.gormDB.WithContext(ctx).Model(&model.Product{})

	if option.Search != nil {
		stmt = stmt.Where("name ILIKE ?", "%"+*option.Search+"%")
	}

	result := stmt.Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *productRepository) Fetch(ctx context.Context, options ...model.ProductQueryOption) ([]model.Product, error) {
	option := model.ProductQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	var products []model.Product
	stmt := r.gormDB.WithContext(ctx)

	if option.Search != nil {
		stmt = stmt.Where("name ILIKE ?", "%"+*option.Search+"%")
	}

	stmt = applyPagination(stmt, option.Limit, option.Page)

	result := stmt.Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}

	return products, nil
}

func (r *productRepository) Get(ctx context.Context, id string) (*model.Product, error) {
	var product model.Product
	result := r.gormDB.WithContext(ctx).First(&product, "id = ?", id)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &product, nil
}
