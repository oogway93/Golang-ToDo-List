package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main_page(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})

}

func main() {
	router := gin.Default()
	router.SetTrustedProxies([]string{"46.159.233.168"})

	router.GET("/", main_page)
	router.Run(":8080")

}
