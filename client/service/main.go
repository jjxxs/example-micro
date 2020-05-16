package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
	"github.com/vesose/example-micro/client"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	logger.DefaultLogger = zl.NewLogger(logger.WithOutput(output), logger.WithLevel(logger.DebugLevel))

	registry := etcdv3.NewRegistry()
	store := redis.NewStore()

	service := micro.NewService(
		micro.Registry(registry),
		micro.Store(store),
	)
	service.Init()

	client := client.New(api.NewGreeterService("greeter", service.Client()), service.Options().Store)

	client.Interact()
}
