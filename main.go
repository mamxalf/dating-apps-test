package main

import (
	"dating-apps/configs"
	"dating-apps/shared/logger"
)

//go:generate go run github.com/swaggo/swag/cmd/swag init
//go:generate go run github.com/google/wire/cmd/wire

var config *configs.Config

func main() {
	logger.InitLogger()
	// Initialize config
	config = configs.Get()

	// Set desired log level
	logger.SetLogLevel(config)

	http := InitializeService()
	http.SetupAndServe()
}
