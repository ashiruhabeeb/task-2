package app

import (
	"net/http"
	"simple-crud-app/app/models"
	"simple-crud-app/utils"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (a *App) Router() {
	a.Gin = gin.Default()

	a.Gin.POST("/api", func(ctx *gin.Context) {
		var p models.Person
		name := ctx.Query("name")

		if err := ctx.ShouldBindQuery(&p); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		
		p = models.Person{
			Name:      name,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
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

	a.Gin.GET("/api/:name", func(ctx *gin.Context) {
		var p models.Person

		// if err := ctx.ShouldBindQuery(&p); err != nil {
		// 	ctx.AbortWithError(http.StatusBadRequest, err)
		// 	return
		// }

		if err := a.DB.Where("name = ?", ctx.Param("name")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": "ERROR, Person NOT FOUND",
				"error": err.Error()},	
			)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"Response": p})
	})

	a.Gin.PUT("/api/:name", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("name = ?", ctx.Param("name")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var _p models.Person

		if err := ctx.ShouldBindJSON(&_p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := a.DB.Model(&p).Updates(models.Person{Name: _p.Name}).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		ctx.JSON(http.StatusOK, gin.H{"Response": p})
	})

	a.Gin.DELETE("/api/:name", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("name = ?", ctx.Param("name")).First(&p).Error; err != nil {
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
