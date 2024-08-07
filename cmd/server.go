package main

import (
	"echommerce/internal/handlers"
	"echommerce/internal/services"

	"echommerce/pkg/database"

	// "net/http"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("HOST_ADDRESS")
	user := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	db := database.Init(host, user, dbPass, dbName, port)

	userService := services.NewUserService(db)
	userHandler := handlers.NewUserHandler(userService)

	e := echo.New()

	e.POST("/users", userHandler.CreateUser)
	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello World!")
	// })
	e.POST("/login", userHandler.Login)
	e.Logger.Fatal(e.Start(":1233"))
}
