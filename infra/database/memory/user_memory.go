package memory

import (
	"errors"

	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
)

var (
	ErrUserExists = errors.New("Usuário já cadastrado com este nome")
	ErrUserNotFound = errors.New("Usuário não encontrado")
	ErrEmptyList = errors.New("A lista está vazia")
)

type UserMemoryRepository struct {
	users map[string]entity.User
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{ users: make(map[string]entity.User)}
}

func (repo *UserMemoryRepository) Create(user entity.User) (entity.User, error) {
	err := repo.UserExists(user.Name)
	if err {
		return entity.User{}, ErrUserExists
	}
	repo.users[user.ID] = user
	return user, nil
}

func (repo *UserMemoryRepository) Update(user entity.User) (entity.User, error)  {
	repo.users[user.ID] = user
	return user, nil
}

func (repo *UserMemoryRepository) GetById(id string) (entity.User, error) {
	if user, ok := repo.users[id]; ok {
		return user, nil
	}
	return entity.User{}, ErrUserNotFound
}

func (repo *UserMemoryRepository) List() ([]entity.User, error) {
	
	if len(repo.users) <= 0 {
		return []entity.User{}, ErrEmptyList
	}

    usersSlice := make([]entity.User, 0, len(repo.users))
    for _, user := range repo.users {
        usersSlice = append(usersSlice, user)
    }
	
	return usersSlice, nil
}

func (repo *UserMemoryRepository) Delete(id string) error {
	if _, ok := repo.users[id]; ok {
		delete(repo.users, id)
		return nil
	}
	return ErrUserNotFound
}

//VALIDATIONS
func (repo *UserMemoryRepository) UserExists(name string) bool {
    for _, user := range repo.users {
        if user.Name == name {
            return true
        }
    }
    return false
}