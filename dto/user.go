package dto

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Admin    bool

	authenticated bool `gorm:"-"`
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) Authenticated() *User {
	u.authenticated = true

	return u
}

func (u *User) IsAdmin() bool {
	return u.Admin
}
