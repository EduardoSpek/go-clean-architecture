package usecase

import "github.com/eduardospek/go-clean-arquiteture/domain/entity"

type InfoInteractor struct {
	InfoRepository entity.InfoRepository
}

func NewInfoInteractor(inforepository entity.InfoRepository) *InfoInteractor {
	return &InfoInteractor{ InfoRepository: inforepository }
}

func (interactor *InfoInteractor) CreateInfo(info entity.Info) (entity.Info, error) {
	newinfo, err := entity.NewInfo(info)
	if err != nil {
		return entity.Info{}, err
	}
	return interactor.InfoRepository.Create(*newinfo)
}