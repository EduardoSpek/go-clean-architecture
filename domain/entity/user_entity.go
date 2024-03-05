package entity

import (
	"sync"

	"github.com/google/uuid"
)

var (	
	UsersMutex sync.Mutex
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Zap  string `json:"zap"`
}

func NewUser(name string, zap string) *User {
	user := &User{
		ID:   uuid.NewString(),
		Name: name,
		Zap:  zap,
	}	
	return user
}

type UserRepository interface {
	Create(user User) error
	GetById(id string) (User, error)
}