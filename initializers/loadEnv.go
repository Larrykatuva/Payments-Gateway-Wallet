package initializers

import (
	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		ErrorLog.Println("Error while loading the env")
	}
	InfoLog.Println("Loaded env successfully")
}
