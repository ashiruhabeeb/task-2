package app

import (
	"fmt"
	"log"
	"simple-crud-app/app/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB 		*gorm.DB
	Gin	*gin.Engine
}

func (a *App) SetUpDB(dbUrl string) {
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	fmt.Println("🚀--<<Successfully connected to database>>--🚀")

	// Migrate person data structure to database
	db.AutoMigrate(&models.Person{})
	fmt.Println("🚀--<<Successfully migrated Person object to database>>--🚀")

	a.DB = db
}

func (a *App) Router() {
	a.Gin = gin.Default()
}

func (a *App) StartServer(port string) {
	fmt.Println("🚀--<<API IS RUNNNING>>--🚀")
	if port == "" {
		port = "8090"
	}
	a.Gin.Run(port)
}