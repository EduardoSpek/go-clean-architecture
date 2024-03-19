package aggregate

import (
	"github.com/eduardospek/go-clean-architecture/domain/entity"
)

type UserWithInfo struct {
	User entity.User `json:"user"`
	Info entity.InfoOutput `json:"info"`
}

func NewUserWithInfo(user entity.User, info entity.InfoOutput) *UserWithInfo {
	return &UserWithInfo{  User: user, Info: info }
}