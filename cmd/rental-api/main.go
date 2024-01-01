package main

import (
	"log"
	"os"

	"parking-lot/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
