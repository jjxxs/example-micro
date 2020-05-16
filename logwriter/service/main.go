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
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/logwriter"
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
	store := redis.NewStore()

	service := micro.NewService(
		micro.Name("logwriter"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Store(store),
	)

	service.Init()

	if err := micro.RegisterSubscriber("log.*", service.Server(),
		logwriter.New(store)); err != nil {
		panic(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}
}
