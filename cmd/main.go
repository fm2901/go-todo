package main

import (
	"log"

	"github.com/fm2901/go-todo"
	"github.com/fm2901/go-todo/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8888", handlers.InitRoutes()); err != nil {
		log.Fatalf("error: %s", err.Error())
	}

}
