package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/tweakdeveloper/ao3pub2/internal/archive"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.GET("/", handleRoot)
	r.GET("/works/:work", handleSimpleWork)
	r.Run()
}

func handleRoot(c *gin.Context) {
	c.String(http.StatusOK, "howdy!")
}

func handleSimpleWork(c *gin.Context) {
	work, err := archive.GetWorkText(c.Param("work"))
	if err != nil {
		log.Print(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"work":  work,
	})
}
