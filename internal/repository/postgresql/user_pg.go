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
	createUser := pg.db.Table("users").Where(user.Username).
		Attrs(user.ID, user.CreatedAt).FirstOrCreate(&user)
	if createUser.Error != nil {
		return "", createUser.Error
	}
	return user.ID.String(), nil
}

func (pg *UserPg) ExistenceUserName(userId string) error {
	panic("implement me")
}
