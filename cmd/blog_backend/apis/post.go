package apis

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetPost is function for endpoint /api/v1/posts to get Post by ID
func GetPost(c *gin.Context) {
	//var ps := new services.PostService(daos.NewPostDAO())
	s := services.NewPostService(daos.NewPostDAO())
	if post, err := s.Get(c.Param("id")); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, post)
	}
}
