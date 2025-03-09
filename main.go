package main

import (
	"log"

	"coop-gardens-be/config"
	"coop-gardens-be/internal/api/routers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load .env
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	// Group API v1
	apiV1 := e.Group("/api/v1")

	routers.LoginRoutes(apiV1)

	// Kết nối DB
	config.InitDB()

	log.Println("🚀 Server đang chạy tại: http://localhost:8080")
	e.Start(":8080")
}
