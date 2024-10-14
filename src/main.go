package main

import (
	"go-blog-service/src/server"
	"log"
)

func main() {

	app := server.CreateBloggingApp()

	if err := app.Listen(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
