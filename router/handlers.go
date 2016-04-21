package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/audio-challenge/socket"
	"github.com/nikitasmall/audio-challenge/task"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

func messageUploadHandler(c *gin.Context) {
	file, _, err := c.Request.FormFile("blob")
	if err != nil {
		log.Print(err)
	}

	var message string
	if t, err := task.ParseMessage(file); err != nil {
		message = err.Error()
	} else {
		t.DefineTask()
		message = t.RawQuery
	}

	socket.MainHub.SendMessage(socket.TaskAdd, gin.H{"text": message})
	c.JSON(http.StatusOK, gin.H{"message": message})
}
