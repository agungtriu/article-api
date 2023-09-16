package middlewares

import (
	"article-api/models/user/database"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	UserId   int    `json:"userId"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateJWT(user database.User) string {
	claims := &jwtCustomClaims{
		int(user.ID),
		user.Username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, _ := token.SignedString([]byte(os.Getenv("PRIVATE_KEY_JWT")))

	return t
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("PRIVATE_KEY_JWT")), nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
