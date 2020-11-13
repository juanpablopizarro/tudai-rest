package main

import (
	"flag"

	"github.com/juanpablopizarro/tudai-rest/internal/config"
	"github.com/juanpablopizarro/tudai-rest/internal/service/greeter"
)

func main() {
	configFile := flag.String("config", "./config.yaml", "this is the service config")
	flag.Parse()

	conf := config.LoadConfig(*configFile)

	service := greeter.NewGreeter()
	http := greeter.NewHTTPGreeter(service, conf)

	http.Run()
}
