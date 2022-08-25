package repositories

import (
	"course/domain"
	"course/pkg/helpers"

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

func (u *userRepository) Create(user *domain.UserModel) (err domain.HttpError) {
	dbError := u.db.Create(user).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}

func (u *userRepository) VerifyAvailableEmail(email string) (err domain.HttpError) {
	var count int64
	dbError := u.db.Model(domain.UserModel{}).Where("email = ?", email).Count(&count).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}

func (u *userRepository) GetUserByEmail(email string, fields ...string) (user *domain.UserModel, err domain.HttpError) {
	dbError := u.db.Select(fields).Where("email = ?", email).First(&user).Error
	err = helpers.CastDatabaseError(dbError, true)
	return
}
