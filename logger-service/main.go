package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/store"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
)

type Sub struct {
	store store.Store
}

func (s *Sub) Process(ctx context.Context, event *api.Event) error {
	md, _ := metadata.FromContext(ctx)
	logger.Infof("Received event %+v with metadata %+v", event, md)

	newSleep := []byte(fmt.Sprintf("%v", rand.Intn(2000)))

	record := store.Record{
		Key:    "sleep",
		Value:  newSleep,
		Expiry: 0,
	}
	err := s.store.Write(&record)

	if err != nil {
		logger.Infof("error while writing to store: %+v", err)
	}

	return err
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
	store := redis.NewStore()

	service := micro.NewService(
		micro.Name("logWriter"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Store(store),
	)

	service.Init()

	if err := micro.RegisterSubscriber("log.*", service.Server(),
		&Sub{
			store: store,
		}); err != nil {
		panic(err)
	}

	if err := service.Run(); err == nil {
		log.Fatal(err)
	}
}
