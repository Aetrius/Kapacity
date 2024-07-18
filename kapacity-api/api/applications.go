package api

import (
	KapacityAPI "kapacity-api/internal/kapacity"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupApplicationRoutes(applications *gin.RouterGroup, db gorm.DB, environmentIn string) {
	applications.GET("/", func(c *gin.Context) {
		getNodes(c, db, environmentIn)
	})

	applications.GET("/:id", func(c *gin.Context) {
		getNode(c, db)
	})
}

func getApplication(c *gin.Context, db gorm.DB) {
	applicationID := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"message": "Get Application", "applicationID": applicationID})
}

func getApplications(c *gin.Context, db gorm.DB, environmentIn string) {
	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	// status := c.DefaultQuery("status", "started")
	// environmentId := c.DefaultQuery("environmentId", environmentIn)

	// Define additional parameters as needed
	additionalParams := map[string]interface{}{
		// "status":         status,
		// "environment_id": environmentId,
	}

	// Sort by created_at or updated_at, defaulting to created_at
	sortColumn := c.DefaultQuery("sortBy", "")
	orderBy := c.DefaultQuery("orderBy", "")
	getPaginatedData(c, db, &KapacityAPI.Application{}, pageStr, pageSizeStr, additionalParams, sortColumn, orderBy)
}
