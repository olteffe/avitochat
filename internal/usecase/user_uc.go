package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/olteffe/avitochat/internal/models"
	"github.com/olteffe/avitochat/internal/repository"
	"time"
)

type UserUseCase struct {
	repo repository.User
}

func NewUserUseCase(repo repository.User) *UserUseCase {
	return &UserUseCase{repo: repo}
}

// CreateUserUseCase func create new user
func (uc *UserUseCase) CreateUserUseCase(user *models.Users) (string, error) {
	// simple input data validation
	if user.Username == "" || len(user.Username) > 50 {
		return "", errors.New("invalid username")
	}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	return uc.repo.CreateUserRepository(user)
}
