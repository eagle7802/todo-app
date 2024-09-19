package main

import (
	"github.com/eagle7802/todo-app"
	handler "github.com/eagle7802/todo-app/pkg/handlers"
	"github.com/eagle7802/todo-app/pkg/repository"
	"github.com/eagle7802/todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	//handlers := new(handler.Handler)
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)

	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
