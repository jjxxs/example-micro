package main

import (
	"log"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	nats "github.com/micro/go-plugins/broker/nats/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/vesose/example-micro/logwriter"
	"github.com/vesose/example-micro/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
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
