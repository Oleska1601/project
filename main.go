package main

import (
	"project/internal/controller"
)

func main() {
	webDir := "./front"
	s := controller.New(webDir)
	s.Run(":8080")
}
