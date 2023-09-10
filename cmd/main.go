package main

import (
	"log"
	"os"

	"github.com/victorvelo/notes/internal/config"
	"github.com/victorvelo/notes/pkg/handler"
	"github.com/victorvelo/notes/pkg/repository"
	"github.com/victorvelo/notes/pkg/server"
	"github.com/victorvelo/notes/pkg/service"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatal("error initializing configs :", err)
	}

	db, err := repository.NewPostgresDB(&config.Postgres)
	if err != nil {
		log.Fatal("failed to initialize db : ", err)
	}

	defer db.Close()

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	server := server.NewServer()

	if err := server.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
