package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

//GetEnv func
func GetEnv(key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalln("Cannot Load .Env File")
	}
	return os.Getenv(key)
}
