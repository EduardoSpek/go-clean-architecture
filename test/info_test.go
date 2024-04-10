package test

import (
	"testing"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	database "github.com/eduardospek/go-clean-architecture/infra/database/memorydb"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
	"github.com/eduardospek/go-clean-architecture/validations"
)

func TestInfo(t *testing.T) {
	t.Parallel()
	user := entity.NewUser("Info Test", "71996229991")

	infoInput := entity.InfoInput{
		Id_user: user.ID,
		Cabelo: "Cacheado",
		Olhos: "Verde",
		Pele: "Branca",
		Corpo: "Magra",

	}
	
	info, err := entity.NewInfo(infoInput)

	if err != nil {
		t.Errorf("Erro: Não foi possível criar a entity Info")
	}


	if info.ID  == "" {
		t.Errorf("Esperado: %s | Recebido: %s", info.ID, "")
	}

	if info.Id_user != infoInput.Id_user {
		t.Errorf("Esperado: %s | Recebido: %s", infoInput.Id_user, info.Id_user)
	}

	if info.Cabelo.String() != infoInput.Cabelo {
		t.Errorf("Esperado: %s | Recebido: %s", infoInput.Cabelo, info.Cabelo.String())
	}

	if info.Olhos.String() != infoInput.Olhos {
		t.Errorf("Esperado: %s | Recebido: %s", infoInput.Olhos, info.Olhos.String())
	}

	if info.Pele.String() != infoInput.Pele {
		t.Errorf("Esperado: %s | Recebido: %s", infoInput.Pele, info.Pele.String())
	}

	if info.Corpo.String() != infoInput.Corpo {
		t.Errorf("Esperado: %s | Recebido: %s", infoInput.Corpo, info.Corpo.String())
	}

	t.Run("Deve retornar erro para tipo de cabelo inválido", func(t *testing.T) {

		infoInput := entity.InfoInput{
			Id_user: user.ID,
			Cabelo: "Encaracolado",
			Olhos: "Verde",
			Pele: "Branca",
			Corpo: "Magra",
	
		}

		esperado := "tipo de cabelo inválido: " + infoInput.Cabelo
		
		_, err := entity.NewInfo(infoInput)
	
		if err.Error() != esperado {
			t.Errorf("Esperado: %s | Recebido: %s", esperado, err.Error())
		}
		

	})

	t.Run("Deve retornar erro para a cor dos solhos inválida", func(t *testing.T) {

		infoInput := entity.InfoInput{
			Id_user: user.ID,
			Cabelo: "Cacheado",
			Olhos: "Vermelho",
			Pele: "Branca",
			Corpo: "Magra",
	
		}

		esperado := "cor dos olhos inválida: " + infoInput.Olhos
		
		_, err := entity.NewInfo(infoInput)
	
		if err.Error() != esperado {
			t.Errorf("Esperado: %s | Recebido: %s", esperado, err.Error())
		}
		

	})

	t.Run("Deve retornar erro para a cor de pele inválida", func(t *testing.T) {

		infoInput := entity.InfoInput{
			Id_user: user.ID,
			Cabelo: "Cacheado",
			Olhos: "Preto",
			Pele: "Amarela",
			Corpo: "Magra",
	
		}

		esperado := "cor de pele inválida: " + infoInput.Pele
		
		_, err := entity.NewInfo(infoInput)
	
		if err.Error() != esperado {
			t.Errorf("Esperado: %s | Recebido: %s", esperado, err.Error())
		}
		

	})

	t.Run("Deve retornar erro para o tipo de corpo inválido", func(t *testing.T) {

		infoInput := entity.InfoInput{
			Id_user: user.ID,
			Cabelo: "Cacheado",
			Olhos: "Castanho",
			Pele: "Parda",
			Corpo: "Obesa",
	
		}

		esperado := "tipo de corpo inválido: " + infoInput.Corpo
		
		_, err := entity.NewInfo(infoInput)
	
		if err.Error() != esperado {
			t.Errorf("Esperado: %s | Recebido: %s", esperado, err.Error())
		}
		

	})

}

func TestInfoInteractor(t *testing.T) {
	t.Parallel()
	userRepo := database.NewUserMemoryRepository()		
	userValidation := validations.NewUserValidation(userRepo)
	userInteractor := usecase.NewUserInteractor(userRepo, *userValidation)

	infoRepo := database.NewInfoMemoryRepository()	
	infoValidation := validations.NewInfoValidation(infoRepo, userRepo)			
	infoInteractor := usecase.NewInfoInteractor(infoRepo, *infoValidation, *userValidation)

	inputUser := entity.User{
		Name: "Eduardo Spek",
		Zap: "71996229991",
	}
	
	user, err := userInteractor.CreateNewUser(inputUser)

	if err != nil {
		t.Errorf("Erro: %s", err)
	}

	infoInput := entity.InfoInput{
		Id_user: user.ID,
		Cabelo: "Cacheado",
		Olhos: "Verde",
		Pele: "Branca",
		Corpo: "Magra",

	}

	t.Run("Deve criar uma nova Informação", func(t *testing.T) {
	
		info, err := infoInteractor.CreateInfo(infoInput)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		if info.Id_user != infoInput.Id_user {
			t.Errorf("Esperado: %s | Recebido: %s", infoInput.Id_user, info.Id_user)
		}
	})	

	t.Run("Deve atualizar uma Informação", func(t *testing.T) {	
		
		inputUser := entity.User{
			Name: "Thaís Freire",
			Zap: "71996229991",
		}
		
		user, err := userInteractor.CreateNewUser(inputUser)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		infoInput := entity.InfoInput{
			Id_user: user.ID,
			Cabelo: "Cacheado",
			Olhos: "Verde",
			Pele: "Branca",
			Corpo: "Magra",

		}
	
		info, err := infoInteractor.CreateInfo(infoInput)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		infoInput = entity.InfoInput{
			ID: info.ID,
			Id_user: user.ID,
			Cabelo: "Liso",
			Olhos: "Azul",
			Pele: "Branca",
			Corpo: "Gorda",
	
		}

		newinfo, err := infoInteractor.UpdateInfo(infoInput)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		if newinfo.Id_user != infoInput.Id_user {
			t.Errorf("Esperado: %s | Recebido: %s", infoInput.Id_user, info.Id_user)
		}

		if newinfo.Cabelo != infoInput.Cabelo {
			t.Errorf("Esperado: %s | Recebido: %s", infoInput.Cabelo, info.Cabelo)
		}

		if newinfo.Olhos != infoInput.Olhos {
			t.Errorf("Esperado: %s | Recebido: %s", infoInput.Olhos, info.Olhos)
		}

		if newinfo.Pele != infoInput.Pele {
			t.Errorf("Esperado: %s | Recebido: %s", infoInput.Pele, info.Pele)
		}

		if newinfo.Corpo != infoInput.Corpo {
			t.Errorf("Esperado: %s | Recebido: %s", infoInput.Corpo, info.Corpo)
		}
	})

}