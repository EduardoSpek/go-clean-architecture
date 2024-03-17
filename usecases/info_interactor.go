package usecase

import "github.com/eduardospek/go-clean-architecture/domain/entity"

type InfoInteractor struct {
	InfoRepository entity.InfoRepository
}

func NewInfoInteractor(inforepository entity.InfoRepository) *InfoInteractor {
	return &InfoInteractor{ InfoRepository: inforepository }
}

func (interactor *InfoInteractor) CreateInfo(info entity.InfoInput) (entity.InfoOutput, error) {
	newinfo, err := entity.NewInfo(info)
	if err != nil {
		return entity.InfoOutput{}, err
	}
	return interactor.InfoRepository.Create(*newinfo)
}