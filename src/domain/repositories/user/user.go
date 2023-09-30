package user

import (
	entities2 "bm-users/src/entities"
	"errors"
)

var ErrUserNotfound = errors.New("user not found")

// UserRepository defines the interface for interacting with the User entity.
type UserRepository interface {
	Create(user *entities2.User) error
	GetByID(id int64) (*entities2.User, error)
	GetAll() ([]*entities2.User, error)
	Update(user *entities2.User) error
	Delete(id int) error
}
