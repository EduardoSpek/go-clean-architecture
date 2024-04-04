package usecase

import (
	"github.com/eduardospek/go-clean-architecture/domain/aggregate"
	"github.com/eduardospek/go-clean-architecture/domain/entity"
)

type InfoRepo interface {
	GetById(id string) (entity.InfoOutput, error)
	GetByIdUser(id string) (entity.InfoOutput, error)
	Create(info entity.Info) (entity.InfoOutput, error)
}

type UserInfoInteractor struct {
	UserRepository UserRepository
	InfoRepository InfoRepo	
}

func NewUserInfoInteractor(user_repository UserRepository, inforepository InfoRepo) *UserInfoInteractor {

	return &UserInfoInteractor{InfoRepository: inforepository, UserRepository: user_repository}
}

func (interactor *UserInfoInteractor) Get(id_user string) (aggregate.UserWithInfo, error) {	
	user, err := interactor.UserRepository.GetById(id_user)
	if err != nil {
		return aggregate.UserWithInfo{}, err
	}	

	info, _ := interactor.InfoRepository.GetByIdUser(id_user)

	userinfo := aggregate.NewUserWithInfo(user, info)	

	return *userinfo, nil
}