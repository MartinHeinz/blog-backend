package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetPost is function for endpoint /api/v1/posts to get Post by ID
func GetPost(c *gin.Context) {
	s := services.NewPostService(daos.NewPostDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if post, err := s.Get(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, post)
	}

}
