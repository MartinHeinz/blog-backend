package apis

import (
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// Add subscriber to mailing list at /api/v1/newsletter/subscribe/
func AddSubscriber(c *gin.Context) {
	url := "https://api.mailerlite.com/api/v2/groups/106705303/subscribers"

	req, err := http.NewRequest("POST", url, c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-MailerLite-ApiKey", config.Config.APIKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		log.Println("Status 401: Check API key in environment variables.")
		c.JSON(http.StatusUnauthorized, gin.H{})
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(fmt.Sprintf("New subscriber added: %s", string(body)))
		c.JSON(resp.StatusCode, gin.H{})
	}
}
