package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/juanpablopizarro/tudai-rest/internal/config"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(cfg.DB.Driver)
	fmt.Println(cfg.Version)

}
