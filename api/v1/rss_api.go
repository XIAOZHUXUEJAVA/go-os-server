package v1

import (
	"fmt"
	"net/http"
	"oschina/model"

	"github.com/gin-gonic/gin"
)

func GetAllRss(c *gin.Context) {
	if rss, err := model.GetAllRss(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		for _, r := range rss {
			fmt.Printf("Title: %s\n", r.Title)
			fmt.Printf("Description: %s\n", r.Description)
			fmt.Println()
		}
		c.JSON(http.StatusOK, rss)
	}
}
