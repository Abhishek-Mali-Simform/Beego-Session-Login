package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/joho/godotenv"
	"os"
)

type envConfigs struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBName     string
	DBPassword string
	DBAlias    string
}

var EnvConfigs *envConfigs

func LoadEnv() {
	EnvConfigs = initializeEnv()
}

func initializeEnv() (config *envConfigs) {
	err := godotenv.Load()
	if err != nil {
		logs.Error("Error loading .env file")
	}
	var env envConfigs
	env.DBHost = os.Getenv("HOST")
	env.DBPort = os.Getenv("DBPORT")
	env.DBUser = os.Getenv("USER")
	env.DBName = os.Getenv("NAME")
	env.DBPassword = os.Getenv("PASSWORD")
	env.DBAlias = os.Getenv("ALIAS")
	config = &env
	return
}
