package apis

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPost is function for endpoint /api/v1/posts to get Post by ID
func GetPost(c *gin.Context) {
	var post models.Post // TODO move DB query to DAOs or Services
	if err := config.Config.DB.Where("id = ?", c.Param("id")).First(&post).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, post)
	}
}
