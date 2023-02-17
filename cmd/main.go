package main

import (
	"log"

	"github.com/fm2901/go-todo"
	"github.com/fm2901/go-todo/pkg/handler"
	"github.com/fm2901/go-todo/pkg/repository"
	"github.com/fm2901/go-todo/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
