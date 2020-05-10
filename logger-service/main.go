package main

import (
	"context"
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/vesose/example-micro/api"
)

type Sub struct{}

//nolint:unparam
func (*Sub) Process(ctx context.Context, event *api.Event) error {
	md, _ := metadata.FromContext(ctx)
	log.Printf("[logger] Received event %+v with metadata %+v\n", event, md)

	return nil
}

func main() {
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

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
