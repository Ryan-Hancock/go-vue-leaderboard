package main

import (
	"net/http"

	"github.com/Ryan-Hancock/go-vue-leaderboard/context"
	"github.com/Ryan-Hancock/go-vue-leaderboard/models"
	"github.com/Ryan-Hancock/go-vue-leaderboard/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	context.Init()
	storage.Init()
	models.Init()

	scoreModel := models.Score{}

	router := gin.Default()
	// Endpoints for the API and Vue client

	router.Use(static.Serve("/", static.LocalFile("web/dist/", true)))
	router.Use(cors.Default())

	api := router.Group("/api")
	{
		api.GET("/scores", func(c *gin.Context) {
			scores := scoreModel.GetAll()
			c.JSON(http.StatusOK, gin.H{
				"data": scores,
			})
		})
		api.POST("/scores", func(c *gin.Context) {
			var json models.Score
			if err := c.ShouldBindJSON(&json); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			scoreModel.NewOrUpdate(json)
			c.JSON(http.StatusOK, gin.H{"status": "inserted"})
		})
		api.PATCH("/score/:name", func(c *gin.Context) {
			type Reason struct {
				LastReason string `json:"reason"`
			}
			var reason Reason
			if err := c.ShouldBindJSON(&reason); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			name := c.Param("name")

			scoreModel.Increase(name, reason.LastReason)
			c.JSON(http.StatusOK, gin.H{"status": "increased"})
		})
	}

	router.Run(":4505")
}
