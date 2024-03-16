package main

import (
	"log"

	database "github.com/eduardospek/go-clean-arquiteture/infra/database/sqlite"
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
	//userRepo := database.NewUserMysqlRepository()
	userRepo := database.NewUserSQLiteRepository()	
	userInteractor := usecase.NewUserInteractor(userRepo)	
	userController := controllers.NewUserController(*userInteractor)

	infoRepo := database.NewInfoSQLiteRepository()	
	infoInteractor := usecase.NewInfoInteractor(infoRepo)	
	infoController := controllers.NewInfoController(*infoInteractor)

	router := routes.NewRouter()
	router.UserRouter(userController)
	router.InfoRouter(infoController)
	
	router.Start(":8080")

}