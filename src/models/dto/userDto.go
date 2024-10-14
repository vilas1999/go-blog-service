package dto

import "go-blog-service/src/models"

type UserRequestDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponseDto struct {
	UserId uint64 `json:"userId"`
	Email  string `json:"email"`
	Token  string `json:"token"`
}

func UserDtoMapper(user models.User, token string) UserResponseDto {
	return UserResponseDto{
		UserId: user.ID,
		Email:  user.EmailId,
		Token:  token,
	}
}
