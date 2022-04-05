package main

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/apis"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/middleware"
	brotli "github.com/anargu/gin-brotli"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	dbprom "gorm.io/plugin/prometheus"
	"log"
)

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(
		prometheus.DefaultGatherer,
		promhttp.HandlerOpts{
			// Opt into OpenMetrics to support exemplars.
			EnableOpenMetrics: true,
		},
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	// load application configurations
	if err := config.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}

	// Creates a router without any middleware by default
	r := gin.New()
	http := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	r.Use(middleware.CORSMiddleware())
	r.Use(brotli.Brotli(brotli.DefaultCompression))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/posts/by-id/:id", apis.GetPost)
		v1.GET("/posts/by-tag/:tag_name", apis.GetPostsByTag)
		v1.GET("/posts/", apis.GetPosts)
		v1.GET("/tags/:post_id", apis.GetTags)
		v1.GET("/sections/:post_id", apis.GetSections)
		v1.GET("/books/", apis.GetBooks)
		v1.GET("/projects/", apis.GetProjects)
		v1.POST("/newsletter/subscribe/", apis.AddSubscriber)
	}

	http.GET("/metrics/", prometheusHandler())

	config.Config.DB, config.Config.DBErr = gorm.Open(postgres.Open(config.Config.DSN))
	if config.Config.DBErr != nil {
		panic(config.Config.DBErr)
	}

	config.Config.DB.Use(dbprom.New(dbprom.Config{
		DBName:          "blog",
		RefreshInterval: 15,
		StartServer:     false,
		MetricsCollector: []dbprom.MetricsCollector{
			&dbprom.Postgres{},
		},
	}))

	// config.Config.DB.Migrator().AutoMigrate(&models.Post{}, &models.Project{}, &models.Section{}, &models.Tag{}, &models.Book{}) // This is needed for generation of schema for postgres image.

	sqlDB, err := config.Config.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	defer sqlDB.Close()

	fmt.Println(fmt.Sprintf("Successfully connected to :%v", config.Config.DSN))

	go r.RunTLS(fmt.Sprintf(":%v", config.Config.ServerPort), config.Config.CertFile, config.Config.KeyFile)
	http.Run(":9080")
}
