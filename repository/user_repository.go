package repository

import (
	"PORTOFOLIO-REAL-TIME-LEADERBOARD/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error)
		Register(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) CheckEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, bool, error) {
	var user entity.User

	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entity.User{}, false, nil
		}
		return entity.User{}, false, err
	}
	return user, true, nil
}

func (r *userRepository) Register(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
