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
	return interactor.UserRepository.Create(user)
}

func (interactor *UserInteractor) GetById(id string) (entity.User, error) {
	return interactor.UserRepository.GetById(id)
}