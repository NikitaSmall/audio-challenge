package router

import (
	"github.com/gin-gonic/gin"
)

// CreateRouter returns a pointer to gin.Engine
// that handles all the incoming requests
func CreateRouter() *gin.Engine {
	router := newRouter()

	router.GET("/", indexHandler)
	router.POST("/message", messageUploadHandler)

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
