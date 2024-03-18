package usecase

import (
	"github.com/eduardospek/go-clean-architecture/domain/entity"
	"github.com/eduardospek/go-clean-architecture/validations"
)


type InfoRepository interface {
	Create(info entity.Info) (entity.InfoOutput, error)
}

type InfoInteractor struct {
	InfoRepository InfoRepository
	InfoValidation validations.InfoValidation
	UserValidation validations.UserValidation
}

func NewInfoInteractor(
	inforepository InfoRepository, 
	infovalidation validations.InfoValidation, 
	uservalidation validations.UserValidation) *InfoInteractor {

		return &InfoInteractor{ InfoRepository: inforepository, UserValidation: uservalidation, 
	
	}
}

func (interactor *InfoInteractor) CreateInfo(info entity.InfoInput) (entity.InfoOutput, error) {	
	newinfo, err := entity.NewInfo(info)
	if err != nil {
		return entity.InfoOutput{}, err
	}

	//Valida se existe um usuário antes de inserir as informações
	err = interactor.UserValidation.UserExsits(newinfo.Id_user)
	if err != nil {
		return entity.InfoOutput{}, err
	}

	//Valida se o usuário já tem as informações
	err = interactor.InfoValidation.InfoRepository.InfoExists(newinfo.Id_user)
	if err != nil {
		return entity.InfoOutput{}, err
	}

	return interactor.InfoRepository.Create(*newinfo)
}