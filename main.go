package main

import (
	"log"

	"github.com/audaykumar/spoti-go/config"
	"github.com/audaykumar/spoti-go/server"
)

func main() {
	appConfig := config.New()
	// appConfig.Print()
	server := server.New(appConfig)
	err := server.Start()
	if err != nil {
		log.Fatalln("An error with the web server occurred:", err)
	}
}
