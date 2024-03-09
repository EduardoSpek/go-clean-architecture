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

func (interactor *UserInteractor) Create(user entity.User) (entity.User, error) {
	newuser, err := entity.NewUser(user.Name, user.Zap)
	if err != nil {
		return entity.User{}, err
	}
	return interactor.UserRepository.Create(*newuser)
}

func (interactor *UserInteractor) Update(user entity.User) error {
	_, err := interactor.UserRepository.GetById(user.ID)
	if err != nil {
		return err
	}
	newuser := &entity.User{
		ID: user.ID,
		Name: user.Name,
		Zap: user.Zap,
	}
	return interactor.UserRepository.Update(*newuser)
}


func (interactor *UserInteractor) GetById(id string) (entity.User, error) {
	return interactor.UserRepository.GetById(id)
}

func (interactor *UserInteractor) UserList() ([]entity.User, error) {
	return interactor.UserRepository.List()
}

func (interactor *UserInteractor) Delete(id string) (error) {
	return interactor.UserRepository.Delete(id)
}