package usecase

import (
	"github.com/eduardospek/go-clean-architecture/domain/entity"
	"github.com/eduardospek/go-clean-architecture/validations"
)

type UserRepository interface {
	Create(user entity.User) (entity.User, error)
	Update(user entity.User) (entity.User, error)
	GetById(id string) (entity.User, error)
	List() ([]entity.User, error)
	Delete(id string) (error)
	UserExists(name string) error
}

type UserInteractor struct {
	UserRepository UserRepository
	UserValidation validations.UserValidation
}

func NewUserInteractor(repository UserRepository, validation validations.UserValidation) *UserInteractor {	
	return &UserInteractor{ UserRepository: repository, UserValidation: validation }
}

func (interactor *UserInteractor) CreateNewUser(user entity.User) (entity.User, error) {
	
	//Validação para evitar nome e zap vazios
	err := interactor.UserValidation.IsValid(user.Name, user.Zap)
	
	if err != nil {
		return entity.User{}, err
	}

	newuser := entity.NewUser(user.Name, user.Zap)

	//Validação para evitar nome e zap vazios
	err = interactor.UserValidation.UserNameExsits(user.Name)
	
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