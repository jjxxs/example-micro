package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
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

	uuid, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("error creating uuid: %v", err)
	}

	if err := g.publisher.Publish(context.Background(), &api.Event{
		Id:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}); err != nil {
		fmt.Printf("error while publishing: %v", err)
	}

	return nil
}

func main() {
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
