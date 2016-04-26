/*Package main holds the inital config and startups the whole project.
 */
package main

import (
	"github.com/nikitasmall/audio-challenge/config"
	"github.com/nikitasmall/audio-challenge/router"
	"github.com/nikitasmall/audio-challenge/task"
)

func init() {
	config.InitEnv(".env")
	config.Connect()
}

func main() {
	r := router.CreateRouter()

	// setup and run task processor in goroutine
	taskProcessor := task.NewProcessor()
	go taskProcessor.Start()

	r.Run()
}
