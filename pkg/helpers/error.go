package helpers

import (
	"course/domain"
	"errors"

	"gorm.io/gorm"
)

func CastDatabaseError(err error, escapeNotFound bool) domain.HttpError {
	if err == nil {
		return nil
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if escapeNotFound {
			return nil
		}
		return domain.NewNotFoundError(errors.New("data tidak ditemukan"))
	}

	return domain.NewInternalServerError(errors.New("internal server error"))
}