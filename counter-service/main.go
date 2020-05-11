package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
)

type Counter struct {
	publisher micro.Event
	counters  map[string]int32
}

func (g *Counter) Inc(ctx context.Context, req *api.IncRequest, rsp *api.SumResponse) error {
	name := req.Name
	g.counters[name]++
	rsp.Counter = g.counters[name]
	msg := fmt.Sprintf("Request for %s, new counter = %d", name, g.counters[name])

	logger.Info(msg)

	uuid, err := uuid.NewRandom()

	if err != nil {
		logger.Errorf("error creating uuid: %v", err)
	}

	if err := g.publisher.Publish(context.Background(), &api.Event{
		Id:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}); err != nil {
		logger.Errorf("error while publishing: %v", err)
	}

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
		micro.Name("counter"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Broker(broker),
	)

	service.Init()

	if err := api.RegisterHelloCounterHandler(service.Server(), &Counter{
		publisher: micro.NewEvent("log.counter", service.Client()),
		counters:  make(map[string]int32),
	}); err != nil {
		log.Fatal(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
