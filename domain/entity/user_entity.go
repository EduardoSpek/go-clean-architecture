package entity

import (
	"time"

	"github.com/google/uuid"
)


type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Zap  string `json:"zap"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewUser(name string, zap string) (*User) {
	user := &User{
		ID:   uuid.NewString(),
		Name: name,
		Zap:  zap,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}	
	return user
}