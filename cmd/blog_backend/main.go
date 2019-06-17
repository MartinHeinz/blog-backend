package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
)

const (
	host     = "blog_db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"
	sslmode  = "disable"
)

func main() {
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

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Post{}, &Section{})

	defer db.Close()

	fmt.Println("Successfully connected!")

	r.Run(":1234")
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
