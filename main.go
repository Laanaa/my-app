package main

import "github.com/Laanaa/my-app/internal/router"

func main() {
	router := router.Router()

	router.Run(":8080")
}