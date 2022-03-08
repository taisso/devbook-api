package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	BaseUrl = ""
	Port    = 0
)

func Load() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error: ", err)
	}

	portEnv, err := strconv.Atoi(os.Getenv("APP_PORT"))

	if err != nil {
		Port = 8080
	}
	Port = portEnv

	BaseUrl = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PORT"))

}
