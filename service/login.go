package service

import (
	"github.com/kodah/blog/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService interface {
	LoginUser(email string, password string) *dto.User
}

type loginInformation struct {
	username string
	password string

	db DBService
}

func StaticLoginService() LoginService {
	return &loginInformation{
		username: "user",
		password: "userpassword",
	}
}

func DynamicLoginService(dbService DBService) LoginService {
	return &loginInformation{
		db: dbService,
	}
}

func (info *loginInformation) LoginUser(username string, password string) *dto.User {
	user := &dto.User{}

	// using static logins
	if info.db == nil {
		return user
	}

	err := info.db.Conn().Transaction(func(tx *gorm.DB) error {
		tx.Model(&dto.User{}).Where("username = ?", username).First(user)

		return tx.Error
	})
	if err != nil {
		return user
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user
	}

	return user.Authenticated()
}
