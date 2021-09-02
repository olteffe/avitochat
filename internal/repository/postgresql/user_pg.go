package postgresql

import (
	"errors"
	mError "github.com/olteffe/avitochat/internal/message_error"
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type UserPg struct {
	db *gorm.DB
}

func NewUserPg(db *gorm.DB) *UserPg {
	return &UserPg{db: db}
}

// ExistenceUser check username in database
func (pg *UserPg) ExistenceUser(user *models.Users) error {
	err := pg.db.Table("users").Where("username = ?", user.Username).
		Limit(1).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
		return mError.ErrCantCreateUserDB
	}
	return mError.ErrUserAlreadyUsed
}

// CreateUserRepository create new user
// It can return wrong error: if another create between ExistenceUser and CreateUserRepository
func (pg *UserPg) CreateUserRepository(user *models.Users) (string, error) {
	createUser := pg.db.Table("users").Create(&user)
	if err := createUser.Error; err != nil {
		return "", mError.ErrCantCreateUserDB
	}
	return user.ID.String(), nil
}
