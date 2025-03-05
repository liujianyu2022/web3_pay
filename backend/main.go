package main

import (
	"log"
	"web3_pay_backend/config"
	"web3_pay_backend/wire"
)

func main() {
	err := config.LoadConfigFile("./config/config.json")

	if err != nil {
		log.Fatalf("failed to load configuration: %v", err)
	}

	app := wire.InitailizeApp()

	app.Run(":9000")
}
