package repo

import "server/model"

type AuthInterface interface {
	LoginUser(user model.User) (string, error)
	RegisterUser(user model.RegisterUser) error
}

