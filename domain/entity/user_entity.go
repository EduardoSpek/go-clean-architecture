package entity

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (	
	UsersMutex sync.Mutex
)

type UserRepository interface {
	Create(user User) (User, error)
	Update(user User) error
	GetById(id string) (User, error)
	List() ([]User, error)
	Delete(id string) (error)
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
		return errors.New("Erro: Nome e Whatsapp são necessários")
	}
	if name == "" {
		return errors.New("Erro: Nome é necessário")
	}
	if zap == "" {
		return errors.New("Erro: Whatsapp é necessário")
	}
	if len(zap) < 13 {
		return errors.New("Erro: Whatsapp deve ter 13 digitos (Ex: 71 98888-7777)")
	}

	return nil

}