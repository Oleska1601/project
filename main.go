package main

import (
	"project/config"
	"project/internal/controller"
	"project/pkg/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		logger.Logger.Error(err.Error())
	}

	webDir := "./front"
	s := controller.New(webDir)
	s.Run(cfg)
}
