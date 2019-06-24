package main

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/apis"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/middleware"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	r.Use(middleware.CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/posts/:id", apis.GetPost)
		v1.GET("/tags/:post_id", apis.GetTags)
		v1.GET("/sections/:post_id", apis.GetSections)
	}

	config.Config.DB, config.Config.DBErr = gorm.Open("postgres", config.Config.DSN)
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	config.Config.DB.AutoMigrate(&models.Post{}, &models.Section{}, &models.Tag{})

	defer config.Config.DB.Close()

	fmt.Println("Successfully connected!")

	r.Run(fmt.Sprintf(":%v", config.Config.ServerPort))
}
