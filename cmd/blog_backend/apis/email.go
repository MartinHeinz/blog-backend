package apis

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MartinHeinz/blog-backend/cmd/blog_backend/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Request struct {
	Email        string   `json:"email"`
	Groups       []string `json:"groups"`
	SubscribedAt string   `json:"subscribed_at"`
	Status       string   `json:"status"`
}

type EmailRequestBody struct {
	Email string
}

// curl -X POST https://localhost:8080/api/v1/newsletter/subscribe/ --data '{"email":"xyz@gmail.com"}'  -kL | jq .
// Add subscriber to mailing list at /api/v1/newsletter/subscribe/
func AddSubscriber(c *gin.Context) {
	url := "https://connect.mailerlite.com/api/subscribers"
	var requestBody EmailRequestBody
	c.BindJSON(&requestBody)

	now := time.Now()
	postBody, _ := json.Marshal(Request{
		Email:        requestBody.Email,
		Groups:       []string{"109703561522185958"}, // Test group: 109703561527428844
		SubscribedAt: now.Format("2006-01-02 15:04:05"),
		Status:       "unconfirmed",
	})

	b := bytes.NewBuffer(postBody)
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Config.APIKey))

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
