package memory

import (
	"errors"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
)

type UserMemoryRepository struct {
	users map[string]entity.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{ users: make(map[string]entity.User)}
}

func (repo *UserMemoryRepository) Create(user entity.User) error {
	repo.users[user.ID] = user
	return nil
}

func (repo *UserMemoryRepository) GetById(id string) (entity.User, error) {
	if user, ok := repo.users[id]; ok {
		return user, nil
	}
	return entity.User{}, errors.New("User not found")
}

