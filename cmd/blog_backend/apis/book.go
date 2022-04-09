package apis

import (
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetBooks is function for endpoint /api/v1/books to get all books
func GetBooks(c *gin.Context) {
	_, span := config.Config.Tracer.Start(c.Request.Context(), "GetBooks")
	defer span.End()
	s := services.NewBookService(daos.NewBookDAO())
	if books, err := s.FindAll(); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"books": books,
		})
	}
}
