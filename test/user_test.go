package test

import (
	"testing"
	"time"

	"github.com/eduardospek/go-clean-architecture/domain/entity"
	database "github.com/eduardospek/go-clean-architecture/infra/database/memorydb"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
	"github.com/eduardospek/go-clean-architecture/validations"
	"github.com/google/uuid"
)

func TestUser(t *testing.T) {
	t.Parallel()
	testuser := struct {
		ID string
		Name string 
		Zap string
		CreatedAt  time.Time
		UpdatedAt  time.Time
	}{
		ID: uuid.NewString(),
		Name: "Eduardo Spek",
		Zap: "71 99622-9991",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	esperado := entity.NewUser(testuser.Name, testuser.Zap)

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

	t.Run("Testando a data de criação", func(t *testing.T) {
		
		CreatedAt := testuser.CreatedAt
		
		if CreatedAt.Day() != esperado.CreatedAt.Day() {
			t.Errorf("Esperado %d Retornado %d", esperado.CreatedAt.Day(), CreatedAt.Day() )
		}

		if CreatedAt.Month() != esperado.CreatedAt.Month() {
			t.Errorf("Esperado %s Retornado %s", esperado.CreatedAt.Month(), CreatedAt.Month() )
		}

		if CreatedAt.Year() != esperado.CreatedAt.Year() {
			t.Errorf("Esperado %d Retornado %d", esperado.CreatedAt.Year(), CreatedAt.Year() )
		}
	})

	t.Run("Testando a data de atualização", func(t *testing.T) {
		
		UpdatedAt := testuser.UpdatedAt
		
		if UpdatedAt.Day() != esperado.UpdatedAt.Day() {
			t.Errorf("Esperado %d Retornado %d", esperado.UpdatedAt.Day(), UpdatedAt.Day() )
		}

		if UpdatedAt.Month() != esperado.UpdatedAt.Month() {
			t.Errorf("Esperado %s Retornado %s", esperado.UpdatedAt.Month(), UpdatedAt.Month() )
		}

		if UpdatedAt.Year() != esperado.UpdatedAt.Year() {
			t.Errorf("Esperado %d Retornado %d", esperado.UpdatedAt.Year(), UpdatedAt.Year() )
		}
	})

}

func TestUserValidation(t *testing.T) {
	t.Parallel()
	
	userRepo := database.NewUserMemoryRepository()
	userValidation := validations.NewUserValidation(userRepo)
	userInteractor := usecase.NewUserInteractor(userRepo, *userValidation)	

	testuser := &entity.User{
		Name: "Creuza", 
		Zap: "71996229991",
	}

	t.Run("Verifica se o usuário é válido", func(t *testing.T) {
		err := userValidation.IsValid(testuser.Name, testuser.Zap)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}
	})	

	t.Run("Verifica se o usuário já existe pelo Name", func(t *testing.T) {
		err := userValidation.UserNameExsits(testuser.Name)

		esperado := "usuário não encontrado"

		if err != nil {
			t.Errorf("Esperado: %s Retornado %s", esperado, err)
		}

	})

	t.Run("Verifica se um usuário existe por ID", func(t *testing.T) {
		user, err := userInteractor.CreateNewUser(*testuser)
		
		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		err = userValidation.UserExsits(user.ID)

		esperado := "usuário não encontrado"

		if err != nil {
			t.Errorf("Esperado: %s Retornado %s", esperado, err)
		}

	})

}


func TestUserInteractor(t *testing.T) {
	t.Parallel()
	
	userRepo := database.NewUserMemoryRepository()
	userValidation := validations.NewUserValidation(userRepo)
	userInteractor := usecase.NewUserInteractor(userRepo, *userValidation)	

	testuser := &entity.User{
		Name: "Briza", 
		Zap: "71996229993",
	}

	t.Run("Cria um novo usuário", func(t *testing.T) {
		_, err := userInteractor.CreateNewUser(*testuser)
		
		if err != nil {
			t.Errorf("Erro: %s", err)
		}
	})

	t.Run("Atualiza dados do usuário", func(t *testing.T) {
		
		testuser := &entity.User{
			Name: "Letícia Spiller", 
			Zap: "71996229993",
		}

		user, err := userInteractor.CreateNewUser(*testuser)
		
		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		
		edituser := entity.User{
			ID: user.ID,
			Name: "Maria", 
			Zap: "71996229992",
		}

		updateuser, err := userInteractor.UpdateUser(edituser)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		if user.Name == updateuser.Name {
			t.Errorf("Esperado: %s Recebido %s", edituser.Name, updateuser.Name)
		}

		if user.Zap == updateuser.Zap {
			t.Errorf("Esperado: %s Recebido %s", edituser.Zap, updateuser.Zap)
		}
	})	

	t.Run("Deleta o usuário", func(t *testing.T) {
		
		testuser := &entity.User{
			Name: "Aline Mineiro", 
			Zap: "71996229995",
		}

		user, err := userInteractor.CreateNewUser(*testuser)
		
		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		err = userInteractor.DeleteUser(user.ID)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}
	})	

	t.Run("Recebe a lista de usuários", func(t *testing.T) {
		
		testuser := &entity.User{
			Name: "Aline Mineiro", 
			Zap: "71996229995",
		}

		_, err := userInteractor.CreateNewUser(*testuser)
		
		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		list, err := userInteractor.UserList()

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		if list == nil {
			t.Errorf("Erro: %s", err)
		}
	})

	t.Run("Retorna o usuário pelo ID", func(t *testing.T) {
		
		testuser := &entity.User{
			Name: "Alane Dias", 
			Zap: "71996229996",
		}

		newuser, err := userInteractor.CreateNewUser(*testuser)
		
		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		user, err := userInteractor.GetById(newuser.ID)

		if err != nil {
			t.Errorf("Erro: %s", err)
		}

		if len(user.ID) < 36 {
			t.Errorf("Erro: %s", err)
		}
	})

}