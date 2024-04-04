package test

import (
	"errors"
	"testing"

	"github.com/eduardospek/go-clean-architecture/domain/aggregate"
	"github.com/eduardospek/go-clean-architecture/domain/entity"
	database "github.com/eduardospek/go-clean-architecture/infra/database/memorydb"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
	"github.com/eduardospek/go-clean-architecture/validations"
)

type TestCase struct {
	Esperado string
	Recebido string
}

func Resultado(t *testing.T, esperado string, recebido string) {
    t.Helper()
    if esperado != recebido {
        t.Errorf("Esperado: %s | Recebido: %s", esperado, recebido)
    }
}

func TestUserInfo(t *testing.T) {
	t.Parallel()

	userRepo := database.NewUserMemoryRepository()
	infoRepo := database.NewInfoMemoryRepository()	
	userinfoInteractor := usecase.NewUserInfoInteractor(userRepo, infoRepo)	

	
	user := entity.NewUser("Eduardo Spek", "71996229991")

	if user.ID == "" {
		t.Error("User sem ID")
	}
	
	infoInput := entity.InfoInput{
		Id_user: user.ID,
		Cabelo: "Crespo",
		Olhos: "Preto",
		Pele: "Negra",
		Corpo: "Atletica",
	}	

	newinfo, err := entity.NewInfo(infoInput)
	if err != nil {
		t.Errorf("erro NewInfo: %s", err)
	}

	infodb, err := userinfoInteractor.InfoRepository.Create(*newinfo)		
	
	if err != nil {
		t.Errorf("erro NewInfoDB: %s", err)
	}	

	if newinfo.ID == "" {
		t.Error("Info sem ID")
	}	

	info, err := userinfoInteractor.InfoRepository.GetById(infodb.ID)	

	if err != nil {
		t.Errorf("erro GetById: %s", err)
	}

	userinfo := aggregate.NewUserWithInfo(*user, info)

	if userinfo.Info.ID != newinfo.ID {
		t.Errorf("Esperado: %s | Recebido: %s", newinfo.ID, userinfo.Info.ID)
	}
}

func TestUserInfoInteractor(t *testing.T) {
	t.Parallel()

	userRepo := database.NewUserMemoryRepository()		
	userValidation := validations.NewUserValidation(userRepo)
	userInteractor := usecase.NewUserInteractor(userRepo, *userValidation)			

	infoRepo := database.NewInfoMemoryRepository()	
	infoValidation := validations.NewInfoValidation(infoRepo, userRepo)	
	//infoRepo := database.NewInfoMysqlRepository()	
	infoInteractor := usecase.NewInfoInteractor(infoRepo, *infoValidation, *userValidation)		

	userinfoInteractor := usecase.NewUserInfoInteractor(userRepo, infoRepo)		

	newuser := entity.User{
		Name: "Eduardo Spek",
		Zap: "719996229991",
	}
	user, err := userInteractor.CreateNewUser(newuser)

	if err != nil {
		t.Errorf("erro NewInfoDB: %s", err)
	}
	
	infoInput := entity.InfoInput{
		Id_user: user.ID,
		Cabelo: "Crespo",
		Olhos: "Preto",
		Pele: "Negra",
		Corpo: "Atletica",
	}	

	newinfo, err := infoInteractor.CreateInfo(infoInput)
	
	if err != nil {
		t.Errorf("erro NewInfoDB: %s", err)
	}

	_, err = userinfoInteractor.Get("id-qualquer-para-gerar-erro")	

	if err != nil {
		
		esperado := errors.New("usuário não encontrado")
		
		if err.Error() != esperado.Error() {
			t.Errorf("Esperado: %s | Recebido: %s", err, esperado)
		}

	}

	userinfo, err := userinfoInteractor.Get(newinfo.Id_user)	

	if err != nil {
		t.Errorf("erro Get: %s", err)
	}

	infotests := []TestCase {
		{
			Esperado: userinfo.Info.ID,
			Recebido: newinfo.ID,
		},
		{
			Esperado: userinfo.User.ID,
			Recebido: user.ID,
		},
		{
			Esperado: userinfo.User.Name,
			Recebido: user.Name,
		},
		{
			Esperado: userinfo.User.Zap,
			Recebido: user.Zap,
		},
		{
			Esperado: userinfo.Info.Cabelo,
			Recebido: newinfo.Cabelo,
		},
		{
			Esperado: userinfo.Info.Olhos,
			Recebido: newinfo.Olhos,
		},
		{
			Esperado: userinfo.Info.Pele,
			Recebido: newinfo.Pele,
		},
		{
			Esperado: userinfo.Info.Corpo,
			Recebido: newinfo.Corpo,
		},
		{
			Esperado: userinfo.Info.CreatedAt.String(),
			Recebido: newinfo.CreatedAt.String(),
		},
		{
			Esperado: userinfo.Info.UpdatedAt.String(),
			Recebido: newinfo.UpdatedAt.String(),
		},		
		
	}

	for _, teste := range infotests {
		Resultado(t, teste.Esperado, teste.Recebido)
	}
	
}