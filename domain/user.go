package domain

import "time"

type UserModel struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	NoHp      *string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *UserModel) TableName() string {
	return "users"
}

type UserRepository interface {
	Create(*UserModel) error

	VerifyAvailableEmail(string) int64
}
