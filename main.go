package main

import (
	"os"

	"github.com/nikitasmall/audio-challenge/config"
	"github.com/nikitasmall/audio-challenge/router"
)

func init() {
	config.InitEnv(".env")
}

func main() {
	r := router.CreateRouter()
	r.Run(os.Getenv("PORT"))
}
