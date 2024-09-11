package main

import (
	"github.com/eagle7802/todo-app"
	handler "github.com/eagle7802/todo-app/pkg/handlers"
	"github.com/eagle7802/todo-app/pkg/repository"
	"github.com/eagle7802/todo-app/pkg/service"
	"log"
)

func main() {
	//handlers := new(handler.Handler)
	repos := repository.NewRepository()
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
