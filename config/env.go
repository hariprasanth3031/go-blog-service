package config

import (
	"os"
)

// Env : Variable to hold the env details
var Env *envData

type envData struct {
	Dbconn string
}

// InitializeEnv : Initializes the environment with required variables
func InitializeEnv() {

	Env = &envData{
		Dbconn: os.Getenv("Dbconn"),
	}

}
