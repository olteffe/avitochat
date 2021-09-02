package usecase

import (
	"github.com/google/uuid"
	mError "github.com/olteffe/avitochat/internal/message_error"
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
		return "", mError.ErrUserInvalid
	}
	// check username in db
	err := uc.repo.ExistenceUser(user)
	if err != nil {
		return "", err
	}
	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	return uc.repo.CreateUserRepository(user)
}
