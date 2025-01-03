package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var DB_HOST string = os.Getenv("DB_HOST")
var DB_USER string = os.Getenv("DB_USER")
var DB_PASSWORD string = os.Getenv("DB_PASSWORD")
var DB_NAME string = os.Getenv("DB_NAME")
var DB_PORT string = os.Getenv("DB_PORT")
var DB_SSLMODE string = os.Getenv("DB_SSLMODE")
var TimeZone string = os.Getenv("TimeZone")
