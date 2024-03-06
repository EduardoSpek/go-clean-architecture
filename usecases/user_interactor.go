package usecase

import (
	"github.com/eduardospek/go-clean-arquiteture/domain/entity"
)

type UserInteractor struct {
	UserRepository entity.UserRepository
}

func NewUserInteractor(repository entity.UserRepository) *UserInteractor {
	return &UserInteractor{ UserRepository: repository}
}

func (interactor *UserInteractor) Create(user entity.User) error {
	newuser := entity.NewUser(user.Name, user.Zap)
	return interactor.UserRepository.Create(*newuser)
}

func (interactor *UserInteractor) GetById(id string) (entity.User, error) {
	return interactor.UserRepository.GetById(id)
}

func (interactor *UserInteractor) UserList() ([]entity.User, error) {
	return interactor.UserRepository.List()
}