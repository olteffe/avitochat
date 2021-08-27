package usecase

import (
	"github.com/olteffe/avitochat/internal/repository"
)

type UserUseCase struct {
	repo repository.User
}

func NewUserUseCase(repo repository.User) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (uc *UserUseCase) CreateUserUseCase() {
	panic("implement me")
}
