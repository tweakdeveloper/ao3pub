package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", handleRoot)
	r.Run()
}

func handleRoot(c *gin.Context) {
	c.String(http.StatusOK, "howdy!")
}
