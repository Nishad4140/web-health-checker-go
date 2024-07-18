package main

import (
	"log"
	"os"

	"github.com/Nishad4140/web-health-checker-go/app"
)

func main() {
	app := app.Start()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
