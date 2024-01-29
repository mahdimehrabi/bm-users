package user

import (
	entities "bm-users/src/entities"
	"errors"
)

var ErrUserNotfound = errors.New("user not found")

// UserRepository defines the interface for interacting with the User entity.
type UserRepository interface {
	Create(user *entities.User) error
	GetByID(id int64) (*entities.User, error)
	GetAll() ([]*entities.User, error)
	Update(user *entities.User) error
	Delete(id int) error
	GetByField(field string, value string) ([]*entities.User, error)
}
