package repositories

import (
	"course/domain"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) domain.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Create(user *domain.UserModel) (err error) {
	err = u.db.Create(user).Error
	return
}

func (u *userRepository) VerifyAvailableEmail(email string) (count int64) {
	u.db.Model(domain.UserModel{}).Where("email = ?", email).Count(&count)
	return
}
