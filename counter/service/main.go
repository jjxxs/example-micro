package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
	"github.com/vesose/example-micro/counter"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	logger.DefaultLogger = zl.NewLogger(logger.WithOutput(output), logger.WithLevel(logger.DebugLevel))

	registry := etcdv3.NewRegistry()
	broker := nats.NewBroker()

	service := micro.NewService(
		micro.Name("counter"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()

	if err := api.RegisterHelloCounterHandler(service.Server(),
		counter.New(micro.NewEvent("log.counter", service.Client()))); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
