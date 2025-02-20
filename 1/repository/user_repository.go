package repository

import (
	"context"
	"myapp/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	// insert
	Insert(ctx context.Context, user *model.User) error
	InsertMany(ctx context.Context, users []model.User) error

	// read
	Count(ctx context.Context) (int, error)
	Get(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	IsExistByEmail(ctx context.Context, email string) (bool, error)
}

type userRepository struct {
	gormDB *gorm.DB
}

func NewUserRepository(
	gormDB *gorm.DB,
) UserRepository {
	return &userRepository{
		gormDB: gormDB,
	}
}

func (r *userRepository) Insert(ctx context.Context, user *model.User) error {
	result := r.gormDB.WithContext(ctx).Create(user)
	return result.Error
}

func (r *userRepository) InsertMany(ctx context.Context, users []model.User) error {
	if len(users) == 0 {
		return nil
	}

	result := r.gormDB.WithContext(ctx).Create(users)
	return result.Error
}

func (r *userRepository) Get(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	result := r.gormDB.WithContext(ctx).First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &user, nil
}

func (r *userRepository) Count(ctx context.Context) (int, error) {
	var count int64
	result := r.gormDB.WithContext(ctx).Model(&model.User{}).Count(&count)

	if result.Error != nil {
		return 0, result.Error
	}

	return int(count), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	result := r.gormDB.WithContext(ctx).First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, returnIfErr(result.Error, gorm.ErrRecordNotFound)
	}

	return &user, nil
}

func (r *userRepository) IsExistByEmail(ctx context.Context, email string) (bool, error) {
	var count int64
	result := r.gormDB.WithContext(ctx).Model(&model.User{}).Where("email = ?", email).Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}
