package main

import (
	"log"

	"github.com/Mycunycu/gofermart/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
