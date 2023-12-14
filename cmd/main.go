package main

import (
	"context"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/martinyonatann/go-unit-test/config"
	"github.com/martinyonatann/go-unit-test/internal/users/delivery/http"
	"github.com/martinyonatann/go-unit-test/internal/users/repository"
	"github.com/martinyonatann/go-unit-test/internal/users/usecase"
	"github.com/martinyonatann/go-unit-test/pkg/databases"
)

func main() {
	cfg, err := config.LoadConfig("config")
	if err != nil {
		panic(err)
	}

	db, err := databases.NewMySQLDB(context.Background(), cfg.Database)
	if err != nil {
		panic(err)
	}

	echoServer := echo.New()

	repository := repository.NewRepository(db)
	usecase := usecase.NewUseCase(repository)
	handlers := http.NewHandlers(usecase)

	http.MapRoutes(echoServer, handlers)

	log.Println(echoServer.Start(fmt.Sprintf(":%s", cfg.Server.Port)))
}
