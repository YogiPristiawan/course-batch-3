package repositories

import (
	"course/domain"
	"errors"

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

func (u *userRepository) GetUserByEmail(email string, fields ...string) (user *domain.UserModel, err error) {
	err = u.db.Select(fields).Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("user belum terdaftar")
			return
		}
	}
	return
}
