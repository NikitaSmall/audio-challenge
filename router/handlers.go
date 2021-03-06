/*Package router holds router and hadnlers for its routes.
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

// Index page handler.
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", nil)
}

// This hadnlers tries to get all the tasks of some type (e.g. pizza tasks)
// and returns them as json array. Returns an error message in case of failure.
func taskListHandler(c *gin.Context) {
	tasks, err := task.List()
	if err == nil {
		c.JSON(http.StatusOK, tasks)
	} else {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
}

// This handler fires up when post request with recorded voice is comming.
// The body should contain audio file (wav) as a binary.
// It is sent immediately to the recognition API. After this work done
// the answer is sent back to client.
func messageUploadHandler(c *gin.Context) {
	messageFile, _, err := c.Request.FormFile("blob")
	if err != nil {
		log.Print(err)
	}
	phone, _ := getSessionPhone(c)
	t, err := task.ProcessMessage(phone, messageFile)

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	} else {
		socket.MainHub.SendMessage(socket.TaskAdd, t)
		c.JSON(http.StatusOK, t)
	}
}
