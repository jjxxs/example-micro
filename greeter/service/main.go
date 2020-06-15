package main

import (
	"time"

	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/vesose/example-micro/api"
	"github.com/vesose/example-micro/greeter"
	"github.com/vesose/example-micro/misc"
)

func main() {
	logger.DefaultLogger = misc.Logger()
	registry := etcdv3.NewRegistry()

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Registry(registry),
		micro.Flags(&cli.IntFlag{
			Name:  "sleep",
			Usage: "sleep some seconds before the startup",
		}),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			sleep := c.Int("sleep")
			if sleep > 0 {
				logger.Infof("sleeping %d seconds before startup", sleep)
				time.Sleep(time.Duration(sleep) * time.Second)
			}
			return nil
		}),
	)

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
