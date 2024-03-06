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

func (repo *UserMemoryRepository) List() ([]entity.User, error) {
	
	if len(repo.users) <= 0 {
		return []entity.User{}, errors.New("The list is null")
	}

    usersSlice := make([]entity.User, 0, len(repo.users))
    for _, user := range repo.users {
        usersSlice = append(usersSlice, user)
    }
	
	return usersSlice, nil
}
