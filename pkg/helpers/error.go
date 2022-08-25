package helpers

import (
	"course/domain"
	"errors"

	"gorm.io/gorm"
)

func CastDatabaseError(err error, throwError bool) domain.HttpError {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if !throwError {
			return nil
		}
		return domain.NewNotFoundError(errors.New("data tidak ditemukan"))
	}

	return domain.NewInternalServerError(err)
}
