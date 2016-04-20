package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/audio-challenge/lib"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func messageUploadHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("blob")
	if err != nil {
		log.Print(err)
	}

	err = lib.SaveMessageFile(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Audio message saved"})
	}
}
