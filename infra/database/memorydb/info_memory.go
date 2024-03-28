package memory

import (
	"errors"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
)

type InfoMemoryRepository struct {
	infos map[string]entity.Info
}

func NewInfoMemoryRepository() *InfoMemoryRepository { 
	return &InfoMemoryRepository{ infos: make(map[string]entity.Info) }
}

func (repo *InfoMemoryRepository) Create(info entity.Info) (entity.InfoOutput, error) {

	repo.infos[info.ID] = info

	newinfo := entity.InfoOutput{
		ID: info.ID,
		Id_user: info.Id_user,
		Cabelo: info.Cabelo.String(),
		Olhos: info.Olhos.String(),
		Pele: info.Pele.String(),
		Corpo: info.Corpo.String(),
	}

	return newinfo, nil
}

func (repo *InfoMemoryRepository) GetById(id string) (entity.InfoOutput, error) {	
	if info, ok := repo.infos[id]; ok {
		newinfo := entity.InfoOutput{
			ID: info.ID,
			Id_user: info.Id_user,
			Cabelo: info.Cabelo.String(),
			Olhos: info.Olhos.String(),
			Pele: info.Pele.String(),
			Corpo: info.Corpo.String(),
		}
		return newinfo, nil
	}
	return entity.InfoOutput{}, ErrUserNotFound
}

//VALIDATIONS
func (repo *InfoMemoryRepository) UserWithInfo(id_user string) error {

	for _, info := range repo.infos {
		if info.Id_user == id_user {
			return errors.New("erro: Usuário já tem informações")
		}
	}
  
    return nil
}