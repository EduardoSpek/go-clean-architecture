package usecase

import "github.com/eduardospek/go-clean-arquiteture/domain/entity"

type InfoInteractor struct {
	InfoRepository entity.InfoRepository
}

func NewInfoInteractor(inforepository entity.InfoRepository) *InfoInteractor {
	return &InfoInteractor{ InfoRepository: inforepository }
}

func (interactor *InfoInteractor) CreateInfo(id_user string, cabelo string, olhos string, pele string, corpo string) (entity.Info, error) {
	newinfo, err := entity.NewInfo(id_user, cabelo, olhos, pele, corpo)
	if err != nil {
		return entity.Info{}, err
	}
	return interactor.InfoRepository.Create(*newinfo)
}