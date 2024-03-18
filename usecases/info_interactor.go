package usecase

import "github.com/eduardospek/go-clean-architecture/domain/entity"


type InfoRepository interface {
	Create(info entity.Info) (entity.InfoOutput, error)
}

type InfoInteractor struct {
	InfoRepository InfoRepository
}

func NewInfoInteractor(inforepository InfoRepository) *InfoInteractor {
	return &InfoInteractor{ InfoRepository: inforepository }
}

func (interactor *InfoInteractor) CreateInfo(info entity.InfoInput) (entity.InfoOutput, error) {
	newinfo, err := entity.NewInfo(info)
	if err != nil {
		return entity.InfoOutput{}, err
	}
	return interactor.InfoRepository.Create(*newinfo)
}