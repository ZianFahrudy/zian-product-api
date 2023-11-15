package repository

import (
	"context"
	"zian-product-api/common/exception"
	"zian-product-api/domain/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(ctx context.Context, user entity.User) entity.User
	FindEmailExist(ctx context.Context, email string) bool
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	FindByID(ctx context.Context, ID int) (entity.User, error)
}

type authRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepositoryImpl(db *gorm.DB) *authRepositoryImpl {
	return &authRepositoryImpl{db}
}

// repositoryimpl method
func (r *authRepositoryImpl) Save(ctx context.Context, user entity.User) entity.User {
	err := r.db.WithContext(ctx).Create(&user).Error
	exception.PanicIfNeeded(err)

	return user
}

func (r *authRepositoryImpl) FindEmailExist(ctx context.Context, email string) bool {
	var user entity.User

	err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error

	if err != nil {
		return false
	}

	return true
}

func (r *authRepositoryImpl) FindByEmail(ctx context.Context, email string) (entity.User, error) {
	var user entity.User

	err := r.db.WithContext(ctx).Where("email = ?", email).Find(&user).Error
	exception.PanicIfNeeded(err)

	return user, nil
}

func (r *authRepositoryImpl) FindByID(ctx context.Context, ID int) (entity.User, error) {
	var user entity.User
	err := r.db.WithContext(ctx).Where("id = ?", ID).Find(&user).Error
	exception.PanicIfNeeded(err)

	return user, nil
}
