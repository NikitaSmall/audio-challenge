/*Package router holds router and hadnlers for its routes.
 * In this file websocket connection handler is declared.
 */
package router

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/nikitasmall/audio-challenge/socket"
)

// basic gorilla/websocket upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// HubHandler is a function handles GET request and
// upgrades it to websocket connection to update connected clients on some event.
func hubHandler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Panic(err)
	}

	client := socket.CreateClient(ws)
	socket.MainHub.Register(client)

	go client.ReadPump()
	client.WritePump()
}
