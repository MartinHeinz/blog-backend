package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// GetPost is function for endpoint /api/v1/posts/by-id/:id to get Post by ID
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

// GetPosts is function for endpoint /api/v1/posts to get all posts
func GetPosts(c *gin.Context) {
	s := services.NewPostService(daos.NewPostDAO())
	if posts, err := s.FindAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"posts": posts,
		})
	}
}

// GetPostsByTag is function for endpoint /posts/by-tag/:tag_name to get all posts with specified tag
func GetPostsByTag(c *gin.Context) {
	s := services.NewPostService(daos.NewPostDAO())
	tag := c.Param("tag_name")
	posts, _ := s.FindAllByTag(tag)
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}
