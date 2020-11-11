package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/juanpablopizarro/tudai-rest/internal/config"
	"github.com/juanpablopizarro/tudai-rest/internal/service/chat"
)

func main() {
	cfg := readConfig()
	service, _ := chat.New(cfg)

	for _, m := range service.FindAll() {
		fmt.Println(m)
	}
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return cfg
}
