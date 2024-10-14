package server

import (
	"fmt"
	"go-blog-service/src/routes"

	"github.com/gofiber/fiber/v2"
)

func CreateBloggingApp() *fiber.App {

	fmt.Println("From the func")

	app := fiber.New()

	postGorup := app.Group("/post")
	userGroup := app.Group("/user")

	routes.PostsRoute(postGorup)
	routes.UserRoute(userGroup)

	return app
}
