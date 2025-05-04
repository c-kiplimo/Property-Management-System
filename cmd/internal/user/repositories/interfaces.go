package repositories

import (
	"tenant-management/cmd/internal/user/domain"
)

type UserRepository interface {
	Save(user domain.User) error
	FindByEmail(email string) (*domain.User, error)
}
