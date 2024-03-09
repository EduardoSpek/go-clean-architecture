package main

import "testing"

type User struct {
	ID string
	Name string
	Zap string
	expected bool
}

func UserExists(users map[string]User, name string) bool {
    for _, user := range users {
        if user.Name == name {
            return true
        }
    }
    return false
}

func TestValidarNome(t *testing.T) {
	users := map[string]User{
        "uuid1": {Name: "Eduardo", Zap: "71 99622-9991"},
        "uuid2": {Name: "Maria", Zap: "71 99622-9991"},
		"uuid3": {Name: "Nathan", Zap: "71 99622-9991"},
    }

	name_test := "Nathan"

	for _, tt := range users {
		t.Run(tt.Name, func(t *testing.T) {
            exists := UserExists(users, name_test)
            if exists {
                t.Errorf("UserExists(%s) = %v, expected %v", name_test, exists, false)
            }
        })
    }	
}