package gorm

import (
	user2 "bm-users/src/domain/repositories/user"
	"bm-users/src/entities"
	"gorm.io/gorm"
)

// GormUserRepository is a GORM-based implementation of the UserRepository interface.
type UserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository creates a new instance of GormUserRepository.
func NewGormUserRepository(db *gorm.DB) *UserRepository {
	db.AutoMigrate(&entities.User{})
	return &UserRepository{db: db}
}

func (repo *UserRepository) GetByField(field string, value string) ([]*entities.User, error) {
	users := make([]*entities.User, 0)
	err := repo.db.Where(field, value).Find(users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create adds a new user to the database.
func (repo *UserRepository) Create(user *entities.User) error {
	return repo.db.Model(user).Create(user).Error
}

// GetByID retrieves a user from the database by its ID.
func (repo *UserRepository) GetByID(id int64) (*entities.User, error) {
	user := &entities.User{}
	if err := repo.db.First(user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, user2.ErrUserNotfound
		}
		return nil, err
	}
	return user, nil
}

// GetAll retrieves all users from the database.
func (repo *UserRepository) GetAll() ([]*entities.User, error) {
	var users []*entities.User
	if err := repo.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update modifies an existing user in the database.
func (repo *UserRepository) Update(user *entities.User) error {
	return repo.db.Save(user).Error
}

// Delete removes an user from the database by its ID.
func (repo *UserRepository) Delete(id int) error {
	return repo.db.Delete(&entities.User{}, id).Error
}
