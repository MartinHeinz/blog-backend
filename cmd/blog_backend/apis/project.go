package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetProjects is function for endpoint /api/v1/projects to get all project
func GetProjects(c *gin.Context) {
	s := services.NewProjectService(daos.NewProjectDAO())
	if projects, err := s.FindAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"projects": projects,
		})
	}
}
