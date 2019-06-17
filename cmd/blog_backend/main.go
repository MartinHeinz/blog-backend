package main

import (
	"fmt"
	"github.com/MartinHeinz/go-vue-blog/cmd/blog_backend/config"
	"github.com/MartinHeinz/go-vue-blog/cmd/blog_backend/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/", testGetData)
		v1.POST("/", testPostData)
	}

	db, err := gorm.Open("postgres", config.Config.DSN)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Post{}, &models.Section{}, &models.Tag{})

	defer db.Close()

	fmt.Println("Successfully connected!")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}

func JSON(c *gin.Context, code int, obj interface{}) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(code, obj)
}

func testGetData(c *gin.Context) {
	JSON(c, http.StatusOK, gin.H{
		"message": "Get Hello",
	})
}

func testPostData(c *gin.Context) {
	JSON(c, http.StatusOK, gin.H{
		"message": "Post Hello",
	})
}
