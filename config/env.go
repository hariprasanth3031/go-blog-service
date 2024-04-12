package config

import (
	"github.com/joho/godotenv"
	"os"
)

// Env : Variable to hold the env details
var Env *envData

type envData struct {
	Dbconn string
}

// InitializeEnv : Initializes the environment with required variables
func InitializeEnv() {

	Logger.Debug("Initializing Env's...")

	_ = godotenv.Load("./config.env")

	Env = &envData{
		Dbconn: os.Getenv("Dbconn"),
	}

	Logger.Debug("Env's successfully initialized!!!")

}
