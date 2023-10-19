package main

import (
	"hw8/config"
	"hw8/internal/app"
)

func main() {
	cfg := config.NewConfig()
	app.Run(cfg)
}
