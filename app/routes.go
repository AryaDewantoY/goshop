package app

import (
	"flag"
	"log"
	"os"

	"github.com/aryadewantoy/goshop/app/controllers"
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func Run() {
	var routes = controllers.Server{}
	var appConfig = controllers.AppConfig{}
	var dbConfig = controllers.DBConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error Running to .env file")
	}

	appConfig.AppName = getEnv("APP_NAME", "GoShopApp")
	appConfig.AppEnv = getEnv("APP_ENV", "Development")
	appConfig.AppPort = getEnv("APP_PORT", "9000")

	dbConfig.DBHost = getEnv("DB_HOST", "localhost")
	dbConfig.DBUser = getEnv("DB_USER", "user")
	dbConfig.DBPassword = getEnv("DB_PASSWORD", "admin")
	dbConfig.DBName = getEnv("DB_NAME", "dbname")
	dbConfig.DBPort = getEnv("DB_PORT", "5432")
	
	flag.Parse()
	arg := flag.Arg(0)

	if arg != "" {
		routes.InitCommands(appConfig, dbConfig)
	} else {
		routes.Initialize(appConfig, dbConfig)
		routes.Run(":" + appConfig.AppPort)
	}
}
