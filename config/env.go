package config

import (
	"github.com/joho/godotenv"
)

// Env : Variable to hold the env details
var Env *envData

type envData struct {
	Dbconn string
}

// InitializeEnv : Initializes the environment with required variables
func InitializeEnv() {

	_ = godotenv.Load("./config.env")

	Env = &envData{
		Dbconn: "hariprasanth:12345@tcp(127.0.0.1:3306)/blogs_management",
	}

}
