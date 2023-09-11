package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	DB 		*gorm.DB
	Gin	*gin.Engine
}

func (a *App) StartServer(port string) {
	fmt.Println("🚀--<<API IS RUNNNING>>--🚀")
	if port == "" {
		port = "8090"
	}
	a.Gin.Run(port)
}