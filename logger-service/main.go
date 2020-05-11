package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
)

type Sub struct{}

//nolint:unparam
func (*Sub) Process(ctx context.Context, event *api.Event) error {
	md, _ := metadata.FromContext(ctx)
	logger.Infof("Received event %+v with metadata %+v\n", event, md)

	return nil
}

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
		micro.Name("logWriter"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()

	if err := micro.RegisterSubscriber("log.*", service.Server(), new(Sub)); err != nil {
		panic(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}
}
