package main

import (
	"log"

	database "github.com/eduardospek/go-clean-architecture/infra/database/sqlite"
	"github.com/eduardospek/go-clean-architecture/interfaces/controllers"
	"github.com/eduardospek/go-clean-architecture/routes"
	usecase "github.com/eduardospek/go-clean-architecture/usecases"
	"github.com/eduardospek/go-clean-architecture/validations"
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

	
	//userRepo := database.NewUserMysqlRepository()
	userRepo := database.NewUserSQLiteRepository()		
	userValidation := validations.NewUserValidation(userRepo)
	userInteractor := usecase.NewUserInteractor(userRepo, *userValidation)		
	userController := controllers.NewUserController(*userInteractor)

	infoRepo := database.NewInfoSQLiteRepository(userRepo)	
	infoValidation := validations.NewInfoValidation(infoRepo, userRepo)	
	//infoRepo := database.NewInfoMysqlRepository()	
	infoInteractor := usecase.NewInfoInteractor(infoRepo, *infoValidation, *userValidation)	
	infoController := controllers.NewInfoController(*infoInteractor)

	userinfoInteractor := usecase.NewUserInfoInteractor(userRepo, infoRepo)	
	userinfoController := controllers.NewUserInfoController(*userinfoInteractor)

	router := routes.NewRouter()
	router.UserRouter(userController)
	router.InfoRouter(infoController)

	router.UserInfoRouter(userinfoController)
	
	router.Start(":8080")

}