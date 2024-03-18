package entity

import (
	"github.com/google/uuid"
)


type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Zap  string `json:"zap"`
}

func NewUser(name string, zap string) (*User) {
	user := &User{
		ID:   uuid.NewString(),
		Name: name,
		Zap:  zap,
	}	
	return user
}