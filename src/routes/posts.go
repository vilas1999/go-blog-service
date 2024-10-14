package routes

import (
	"go-blog-service/src/controller"
	"go-blog-service/src/models"
	"go-blog-service/src/models/dto"
	"go-blog-service/src/security"

	"github.com/gofiber/fiber/v2"
)

var postsController controller.PostsContoller

func PostsRoute(r fiber.Router) {
	postsController = controller.CreatePostsContoller()
	r.Post("/", security.AuthMiddleware, createPost)
	r.Get("/:postid", security.AuthMiddleware, getPost)
}

func createPost(ctx *fiber.Ctx) error {

	user := ctx.Locals("user").(*models.User)

	postDto := new(dto.PostDto)
	ctx.BodyParser(postDto)

	createdPost, err := postsController.AddPost(postDto.Content, postDto.Title, user.ID)
	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(createdPost)
}

func getPost(ctx *fiber.Ctx) error {
	postId := ctx.Params("postid")
	post, err := postsController.GetPost(postId)

	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	return ctx.Status(fiber.StatusFound).JSON(post)
}
