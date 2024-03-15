package main

import (
	"log"

	sqlite "github.com/eduardospek/go-clean-arquiteture/infra/database/mysql"
	"github.com/eduardospek/go-clean-arquiteture/interfaces/controllers"
	"github.com/eduardospek/go-clean-arquiteture/routes"
	usecase "github.com/eduardospek/go-clean-arquiteture/usecases"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }  
}

func main() {

	//Carrega as variáveis de ambiente
	LoadEnv()

	//userRepo := memory.NewUserMemoryRepository()
	userRepo := sqlite.NewUserMysqlRepository()	
	userInteractor := usecase.NewUserInteractor(userRepo)	
	userController := controllers.NewUserController(*userInteractor)

	router := routes.NewRouter()
	router.SetupRouter(userController)
	router.Start(":8080")

}