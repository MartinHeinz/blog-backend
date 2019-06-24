package apis

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetSections is function for endpoint /api/v1/tags to get all sections by post_id
func GetSections(c *gin.Context) {
	s := services.NewSectionService(daos.NewSectionDAO())
	id, _ := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if sections, err := s.FindAll(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"sections": sections,
		})
	}
}
