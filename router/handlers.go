/*
 * This package holds router and hadnlers for its routes.
 * In this file handlers are declared.
 * They serve as basic route actions.
 */
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
	t, err := task.ProcessMessage(file)

	if err != nil {
		message = err.Error()
	} else {
		message = t.Query()
	}

	socket.MainHub.SendMessage(socket.TaskAdd, message)
	c.JSON(http.StatusOK, t)
}
