package usecase

import (
	"github.com/eduardospek/go-clean-architecture/domain/entity"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	GetById(id string) (entity.User, error)
	List() ([]entity.User, error)
	Delete(id string) (error)
	UserExists(name string) bool
}

type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(repository UserRepository) *UserInteractor {	
	return &UserInteractor{ UserRepository: repository}
}

func (interactor *UserInteractor) CreateNewUser(user entity.User) (entity.User, error) {
	newuser, err := entity.NewUser(user.Name, user.Zap)
	if err != nil {
		return entity.User{}, err
	}
	return interactor.UserRepository.Create(*newuser)
}

func (interactor *UserInteractor) UpdateUser(user entity.User) (entity.User, error) {
	_, err := interactor.UserRepository.GetById(user.ID)
	if err != nil {
		return entity.User{}, err
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

func (interactor *UserInteractor) DeleteUser(id string) (error) {
	return interactor.UserRepository.Delete(id)
}