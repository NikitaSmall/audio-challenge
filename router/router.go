/*Package router holds router and hadnlers for its routes.
 * In this file basic router is generated, with basic preferences.
 * Also MainHub (for websocket connection purposes) starts to work here.
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/nikitasmall/audio-challenge/socket"
)

// CreateRouter returns a pointer to gin.Engine
// that handles all the incoming requests
func CreateRouter() *gin.Engine {
	router := newRouter()

	router.GET("/", indexHandler)
	router.GET("/tasks", taskListHandler)

	router.POST("/message", messageUploadHandler)

	router.GET("/user", checkUser)
	router.POST("/register", register)
	router.POST("/login", login)
	router.DELETE("/logout", logout)

	go socket.MainHub.Run()
	router.GET("/socket", hubHandler)

	return router
}

func newRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")

	return r
}
