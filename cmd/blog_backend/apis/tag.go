package apis

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/daos"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// GetTags is function for endpoint /api/v1/tags to get all tags by post_id
func GetTags(c *gin.Context) {
	s := services.NewTagService(daos.NewTagDAO())
	id, _ := strconv.ParseUint(c.Param("post_id"), 10, 32)
	if tags, err := s.FindAll(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"tags": tags,
		})
		//j, err := json.Marshal(gin.H{
		//	"tags": tags,
		//})
		//fmt.Println(string(j), err)
	}
}
