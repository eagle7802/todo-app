package main

import (
	"github.com/eagle7802/todo-app"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}

}
