package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", indexHandler)
	router.POST("/message", messageUploadHandler)

	router.Run(":3000")
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func messageUploadHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("blob")
	out, err := os.Create("./static/message.wav")
	if err != nil {
		log.Print(err)
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
}
