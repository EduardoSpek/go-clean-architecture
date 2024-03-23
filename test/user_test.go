package test

import (
	"testing"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	database "github.com/eduardospek/go-clean-architecture/infra/database/sqlite"
	"github.com/eduardospek/go-clean-architecture/validations"
	"github.com/google/uuid"
)

func TestUser(t *testing.T) {
	testuser := struct {
		ID string
		Name string 
		Zap string
	}{
		ID: uuid.NewString(), Name: "Eduardo Spek", Zap: "71 99622-9991",
	}

	esperado := &entity.User{
		ID: uuid.NewString(),
		Name: "Eduardo Spek",
		Zap: "71 99622-9991",
	}

	t.Run("ID não pode ter menos que 36 caracteres", func(t *testing.T) {
		userID := testuser.ID
		if len(userID) < 36 {
			t.Errorf("Esperado %d Retornado %d", len(esperado.ID), len(userID))
		}
	})

	t.Run("ID não pode ter mais que 36 caracteres", func(t *testing.T) {
		userID := testuser.ID
		if len(userID) > 36 {
			t.Errorf("Esperado %d Retornado %d", len(esperado.ID), len(userID))
		}
	})
	
	t.Run("Testando o name", func(t *testing.T) {
		userName := testuser.Name
		if userName != esperado.Name {
			t.Errorf("Esperado %s Retornado %s", esperado.Name, userName)
		}
	})

	t.Run("Testando o zap", func(t *testing.T) {
		userZap := testuser.Zap
		if userZap != esperado.Zap {
			t.Errorf("Esperado %s Retornado %s", esperado.Zap, userZap)
		}
	})

}

func TestUserValidation(t *testing.T) {
	testuser := struct {
		ID string
		Name string 
		Zap string
	}{
		ID: uuid.NewString(), Name: "Eduardo Spek", Zap: "71996229991",
	}
	
	userRepo := database.NewUserSQLiteRepository()
	userValidation := validations.NewUserValidation(userRepo)

	t.Run("Verifica se o usuário é válido", func(t *testing.T) {
		err := userValidation.IsValid(testuser.Name, testuser.Zap)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}
	})

}