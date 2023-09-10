package main

import (
	"log"
	"simple-crud-app/app"
	"simple-crud-app/config"
)

func main() {
	cfg, err := config.LoadEnvVariables(".")
	if err != nil {
		log.Fatal("unable to load environment variables")
	}
	a := app.App{}

	a.SetUpDB(cfg.Db_URL)

	a.Router()

	a.StartServer(":"+cfg.Server_Port)
}