package api

import (
	"fmt"
	"log"
	"simple-crud-app/api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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