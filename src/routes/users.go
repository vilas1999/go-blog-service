package routes

import (
	"go-blog-service/src/controller"
	"go-blog-service/src/models/dto"
	"go-blog-service/src/security"

	"github.com/gofiber/fiber/v2"
)

var userController controller.UserController

func UserRoute(r fiber.Router) {
	userController = controller.CreateUserController()
	r.Post("/", addUser)
}

func addUser(ctx *fiber.Ctx) error {
	userDto := new(dto.UserRequestDto)
	ctx.BodyParser(userDto)

	user, err := userController.AddUser(userDto.Email, userDto.Password)

	if err != nil {
		return ctx.Status(fiber.ErrBadRequest.Code).JSON(err)
	}

	token := security.CreateJWTFromUser(user)

	return ctx.Status(fiber.StatusCreated).JSON(dto.UserDtoMapper(*user, token))
}
