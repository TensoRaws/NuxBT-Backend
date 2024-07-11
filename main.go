package main

import (
	"log"
	"os"

	"github.com/TensoRaws/NuxBT-Backend/cmd"
)

const version = "v0.0.1"

func main() {
	app := cmd.NewApp()
	app.Name = "NuxBT-Backend"
	app.Usage = "NuxBT Backend Server"
	app.Description = "A simple backend server for NuxBT"
	app.Version = version

	err := app.Run(os.Args)
	if err != nil {
		log.Printf("Failed to run with %s: %v\\n", os.Args, err)
	}
}
