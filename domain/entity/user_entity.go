package entity

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNameAndZapEmpty = errors.New("Erro: Nome e Whatsapp são necessários")
	ErrNameEmpty = errors.New("Erro: Nome é necessário")
	ErrZapEmpty = errors.New("Erro: Whatsapp é necessário")
	ErrZapLimit = errors.New("Erro: Whatsapp deve ter 13 digitos (Ex: 71 98888-7777)")
)

type UserRepository interface {
	Create(user User) (User, error)
	Update(user User) (User, error)
	GetById(id string) (User, error)
	List() ([]User, error)
	Delete(id string) (error)
	UserExists(name string) bool
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Zap  string `json:"zap"`
}

func NewUser(name string, zap string) (*User, error) {
	err := isValid(name, zap)
	
	if err != nil {
		return &User{}, err
	}

	user := &User{
		ID:   uuid.NewString(),
		Name: name,
		Zap:  zap,
	}	
	return user, nil
}

func isValid(name string, zap string) error {
	if name == "" && zap == "" {
		return ErrNameAndZapEmpty
	}
	if name == "" {
		return ErrNameEmpty
	}
	if zap == "" {
		return ErrZapEmpty
	}
	if len(zap) < 13 {
		return ErrZapLimit
	}

	return nil

}