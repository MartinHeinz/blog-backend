package test_data

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"io/ioutil"
	"strings"
)

func init() {
	// the test may be started from the home directory or a subdirectory
	err := config.LoadConfig("/config") // on host use absolute path
	if err != nil {
		panic(err)
	}
	config.Config.DB, config.Config.Err = gorm.Open("sqlite3", ":memory:")
	if config.Config.Err != nil {
		panic(config.Config.Err)
	}

	config.Config.DB.AutoMigrate(&models.Post{}, &models.Section{}, &models.Tag{})
}

func ResetDB() *gorm.DB {
	if err := runSQLFile(config.Config.DB, getSQLFile()); err != nil {
		panic(fmt.Errorf("error while initializing test database: %s", err))
	}
	return config.Config.DB
}

func getSQLFile() string {
	return "/test_data/db.sql" // on host use absolute path
}

func runSQLFile(db *gorm.DB, file string) error {
	s, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(string(s), ";")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if result := db.Exec(line); result.Error != nil {
			fmt.Println(line)
			return result.Error
		}
	}
	return nil
}
