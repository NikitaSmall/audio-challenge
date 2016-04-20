package main

import (
	"os"

	"github.com/nikitasmall/audio-challenge/router"
	"github.com/nikitasmall/master-service/config"
)

func init() {
	config.InitEnv(".env")
}

func main() {
	r := router.CreateRouter()
	r.Run(os.Getenv("PORT"))
}
