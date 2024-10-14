package security

import (
	"fmt"
	"go-blog-service/src/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
)

var JWTKey = []byte("MIIBOQIBAAJAXWRPQyGlEY+SXz8Uslhe+MLjTgWd8lf/nA0hgCm9JFKC1tq1S73c")

func CreateJWTFromUser(u *models.User) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"sub": strconv.Itoa(int(u.ID)),
		"iat": jwt.NewNumericDate(time.Now()),
		"exp": jwt.NewNumericDate(time.Now().AddDate(0, 0, 3)),
	})

	tokenString, err := token.SignedString(JWTKey)

	if err != nil {
		fmt.Println(err)
		panic("Unable to create JWT")
	}

	return tokenString
}

func AuthenticateJWT(token string) (*models.User, error) {
	claims := jwt.MapClaims{}

	_, error := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return JWTKey, nil
	})

	if error != nil {
		return nil, error
	}

	sub := lo.Must(strconv.Atoi(lo.Must(claims.GetSubject())))
	exp := lo.Must(claims.GetExpirationTime())

	if exp.Before(time.Now()) {
		return nil, jwt.ErrTokenExpired
	}

	return &models.User{
		ID: uint64(sub),
	}, nil
}
