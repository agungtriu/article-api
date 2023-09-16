package helper

import (
	"os"
	"regexp"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(e)
}

func HashPassword(password string) (string, error) {
	key, _ := strconv.Atoi(os.Getenv("PRIVATE_KEY_BCRYPT"))
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), key)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err != nil
}

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		return ":8000"
	}
	return ":" + port
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failed load env file")
	}
}
