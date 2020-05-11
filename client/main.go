package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
)

type Client struct {
	greeter api.GreeterService
}

func (c Client) interact() {
	for i := 0; i < 20; i++ {
		name := "Linda"
		if i%3 == 0 {
			name = "John"
		}

		rsp, err := c.greeter.Hello(context.TODO(), &api.HelloRequest{
			Name: name,
		})
		if err != nil {
			logger.Error(err)
		} else {
			logger.Infof("Received: %+v", rsp.GetGreeting())
		}

		logger.Info("sleeping...")
		time.Sleep(1 * time.Second)
		logger.Info("awake")
	}
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
		micro.Registry(registry),
	)
	service.Init()

	// create the greeter client using the service name and client
	client := Client{
		greeter: api.NewGreeterService("greeter", service.Client()),
	}

	client.interact()
}
