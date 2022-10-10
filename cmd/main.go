package main

import (
	"log"
	todo "todo-app"
	"todo-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	log.Println("server start")
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {

		log.Fatalf("error occurred while running http server %v", err.Error())
	}
}
