package main

import (
	"github.com/eagle7802/todo-app"
	handler "github.com/eagle7802/todo-app/pkg/handlers"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
