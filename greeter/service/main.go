package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
	"github.com/vesose/example-micro/greeter"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	logger.DefaultLogger = zl.NewLogger(logger.WithOutput(output), logger.WithLevel(logger.DebugLevel))

	registry := etcdv3.NewRegistry()

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	service.Init()

	counter := micro.NewService()
	counter.Init()

	if err := api.RegisterGreeterHandler(service.Server(),
		greeter.New(api.NewHelloCounterService("counter", counter.Client()))); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
