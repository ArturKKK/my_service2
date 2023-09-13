package store

import "github.com/ArturKKK/my_service2/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
