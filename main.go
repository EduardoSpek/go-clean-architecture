package main

import (
	"github.com/eduardospek/go-clean-arquiteture/infra/database/memory"
	"github.com/eduardospek/go-clean-arquiteture/interfaces/controllers"
	"github.com/eduardospek/go-clean-arquiteture/routes"
	usecase "github.com/eduardospek/go-clean-arquiteture/usecases"
)

func main() {

	userRepo := memory.NewUserMemoryRepository()
	userInteractor := usecase.NewUserInteractor(userRepo)
	userController := controllers.NewUserController(*userInteractor)

	router := routes.NewRouter()
	router.SetupRouter(userController)
	router.Start(":8080")

}