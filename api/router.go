package api

import (
	"net/http"
	"simple-crud-app/api/models"
	"simple-crud-app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a *App) Router() {
	a.Gin = gin.Default()

	a.Gin.POST("/api/create", func(ctx *gin.Context) {
		var p models.Person
		name := ctx.Query("name")
		
		p = models.Person{
			Name:      name,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		}

		if err := ctx.ShouldBindQuery(&p); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}

		err := utils.InputValidator(p)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		
		if err := a.DB.Create(&p).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"Response": "Person Created!"})
	})

	a.Gin.GET("/api/get", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("name = ?", ctx.Query("name")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": "ERROR, Person NOT FOUND",
				"error": err.Error()},	
			)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Response": p})
	})

	a.Gin.PUT("/api/udt", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("name = ?", ctx.Query("name")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var _p models.Person
		_p.Name = ctx.Query("name")

		if err := ctx.ShouldBindQuery(&_p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := a.DB.Model(&p).Updates(_p).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"Response": _p})
	})

	a.Gin.DELETE("/api/del", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("name = ?", ctx.Query("name")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if err := a.DB.Delete(&p).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Response": "Person Deleted"})
	})
}