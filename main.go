package main

import (
	"github.com/Laanaa/my-app/internal/config"
	"github.com/Laanaa/my-app/internal/database"
	"github.com/Laanaa/my-app/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	database.Connect(cfg)
	router := router.Router()

	router.Run(":8080")
}