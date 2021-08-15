package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/olteffe/avitochat/internal/utils"
)

var ConfigPath = "config.json"

// main func application entry point
func main() {
	var config utils.Config
	configFile, err := os.Open(ConfigPath)
	if err != nil {
		log.Fatalf("cannot open config file: %s", err.Error())
	}

	err = json.NewDecoder(configFile).Decode(&config)
	if err != nil {
		log.Fatalf("cannot unmarshal config file: %s", err.Error())
	}
	configFile.Close()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	utils.StartServer(quit, config)
}
