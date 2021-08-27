package postgresql

import (
	"github.com/olteffe/avitochat/internal/models"
	"gorm.io/gorm"
)

type UserPg struct {
	db *gorm.DB
}

func NewUserPg(db *gorm.DB) *UserPg {
	return &UserPg{db: db}
}

func (pg *UserPg) CreateUserRepository(user *models.Users) (string, error) {
	panic("implement me")
}

func (pg *UserPg) ExistenceUserName(userId string) error {
	panic("implement me")
}
