package validations

import (
	"errors"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
)

var (
	ErrNameAndZapEmpty = errors.New("erro: Nome e Whatsapp são necessários")
	ErrNameEmpty = errors.New("erro: Nome é necessário")
	ErrZapEmpty = errors.New("erro: Whatsapp é necessário")
	ErrZapLimit = errors.New("erro: Whatsapp deve ter minímo de 11 e máximo de 13 digitos (Ex: 71999997777 ou 71 99999-8888)")
)

type UserValidationRepository interface {
	GetById(id string) (entity.User, error)
	UserExists(name string) error
}

type UserValidation struct {
	UserRepository UserValidationRepository
}

func NewUserValidation(repository UserValidationRepository) *UserValidation {
	return &UserValidation{ UserRepository: repository }
}

func (v *UserValidation) IsValid(name string, zap string) error {
	if name == "" && zap == "" {
		return ErrNameAndZapEmpty
	}
	if name == "" {
		return ErrNameEmpty
	}
	if zap == "" {
		return ErrZapEmpty
	}
	if len(zap) < 11 {
		return ErrZapLimit
	}
	if len(zap) > 13 {
		return ErrZapLimit
	}

	return nil
}

func (v *UserValidation) UserExsits(id_user string) error {
	_, err := v.UserRepository.GetById(id_user)
	if err != nil {
		return err
	}
	return nil
}
func (v *UserValidation) UserNameExsits(name string) error {
	err := v.UserRepository.UserExists(name)
	if err != nil {
		return err
	}
	return nil
}
