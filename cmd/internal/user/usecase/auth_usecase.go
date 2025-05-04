package usecase

import (
	"errors"
	"tenant-management/cmd/internal/user/domain"
	"tenant-management/cmd/internal/user/repositories"
	"tenant-management/cmd/internal/user/service"
)

type AuthUsecase struct {
	repo       repositories.UserRepository
	jwtService service.JWTService
}

func NewAuthUsecase(repo repositories.UserRepository, jwt service.JWTService) *AuthUsecase {
	return &AuthUsecase{repo: repo, jwtService: jwt}
}

func (u *AuthUsecase) Register(user domain.User, rawPassword string) error {
	existingUser, _ := u.repo.FindByEmail(user.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashed := service.HashPassword(rawPassword)
	user.PasswordHash = hashed
	return u.repo.Save(user)
}

func (u *AuthUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil || !service.CheckPasswordHash(password, user.PasswordHash) {
		return "", errors.New("invalid credentials")
	}
	return u.jwtService.GenerateToken(user), nil
}
