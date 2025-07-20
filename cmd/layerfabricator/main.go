// cmd/layerfabricator/main.go
package main

import (
	"flag"
	"log"
	"os"

	"layerfabricator/internal/layerfabricator"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	app := layerfabricator.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
