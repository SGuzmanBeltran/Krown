package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl string
	SecretKey string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		DBUrl: getEnv("DB_URL"),
		SecretKey: getEnv("SECRET_KEY"),
	}
}

func getEnv(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Fatal("Could get ENV variable", key)
	return ""
}
