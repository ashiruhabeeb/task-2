package api

import (
	"net/http"
	"simple-crud-app/api/models"
	"simple-crud-app/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (a *App) Router() {
	a.Gin = gin.Default()

	r1 := a.Gin.Group("/api")

	r1.POST("", func(ctx *gin.Context) {
		var p models.Person
		name := ctx.Query("name")

		id := uuid.NewString()
		
		p = models.Person{
			ID:		id,
			Name:	name,
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
		ctx.JSON(http.StatusOK, gin.H{"data": p})
	})

	r1.GET("/", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("id = ?", ctx.Query("id")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": "ERROR, Person not found",
				"error": err.Error()},	
			)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": p})
	})

	r1.GET("", func(ctx *gin.Context) {
		var p []models.Person

		a.DB.Find(&p)

		if len(p) == 0 {
			ctx.JSON(http.StatusOK, gin.H{"response": "No record(s) found"})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"data": p})
	})

	r1.PUT("/", func(ctx *gin.Context) {
		var p models.Person
		// person := ctx.Query("name")

		if err := a.DB.First(&p, "id = ?", ctx.Query("id")).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		var updatePersonData models.Payload

		updatePersonData.Name = ctx.Query("name")

		p.Name = updatePersonData.Name

		if err := ctx.ShouldBindQuery(&p); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": true,
				"msg": 	 err.Error(),
			})
			return
		}

		a.DB.Save(&p)

		ctx.JSON(http.StatusOK, gin.H{"data": p})
	})

	r1.DELETE("/", func(ctx *gin.Context) {
		var p models.Person

		if err := a.DB.Where("id = ?", ctx.Query("id")).First(&p).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		if err := a.DB.Delete(&p).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"status": "Person Deleted"})
	})
}