package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
)

type Greeter struct {
	counter api.HelloCounterService
}

func (g *Greeter) Hello(ctx context.Context, req *api.HelloRequest, rsp *api.HelloResponse) error {
	counterRsp, err := g.counter.Inc(context.Background(), &api.IncRequest{
		Name: req.Name,
	})

	if err != nil {
		logger.Error(err)

		rsp.Greeting = "Hiho " + req.Name
	} else {
		logger.Infof("Counter = %d", counterRsp.Counter)
		if counterRsp.GetCounter()%2 == 1 {
			rsp.Greeting = "Hello " + req.Name
		} else {
			rsp.Greeting = "See you " + req.Name
		}
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

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Registry(registry),
	)

	service.Init()

	counter := micro.NewService()
	counter.Init()

	if err := api.RegisterGreeterHandler(service.Server(), &Greeter{
		counter: api.NewHelloCounterService("counter", counter.Client()),
	}); err != nil {
		logger.Fatal(err)
	}

	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}
