package main

import (
	"log"
	"simple-crud-app/api"
	"simple-crud-app/config"
)

func main() {
	cfg, err := config.LoadEnvVariables(".")
	if err != nil {
		log.Fatal("unable to load environment variables")
	}
	
	a := api.App{}

	a.SetUpDB(cfg.External_URL)

	a.Router()

	a.StartServer(":"+cfg.Server_Port)
}