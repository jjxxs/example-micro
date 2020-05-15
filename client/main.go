package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	zl "github.com/micro/go-plugins/logger/zerolog/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
	"github.com/micro/go-plugins/store/redis/v2"
	"github.com/rs/zerolog"
	"github.com/vesose/example-micro/api"
)

type Client struct {
	greeter api.GreeterService
	store   store.Store
}

func (c Client) interact() {
	for i := 0; i < 20; i++ {
		name := "Linda"
		if i%3 == 0 {
			name = "John"
		}

		rsp, err := c.greeter.Hello(context.Background(), &api.HelloRequest{
			Name: name,
		})
		if err != nil {
			logger.Error(err)
		} else {
			logger.Infof("Received: %+v", rsp.GetGreeting())
		}

		res, err := c.store.Read("sleep", func(r *store.ReadOptions) { r.Table = "sleeper" })

		sleep := 1000

		if err != nil {
			logger.Errorf("error while reading from store: %+v", err)
		} else {
			sleep, err = strconv.Atoi(string(res[0].Value))
			if err != nil {
				logger.Errorf("error while converting value from store: %+v", err)
				sleep = 1
			} else {
				logger.Infof("read from store: %+v", sleep)
			}
		}

		logger.Infof("sleeping %v milliseconds...", sleep)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
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
	store := redis.NewStore()

	service := micro.NewService(
		micro.Registry(registry),
		micro.Store(store),
	)
	service.Init()

	// create the greeter client using the service name and client
	client := Client{
		greeter: api.NewGreeterService("greeter", service.Client()),
		store:   service.Options().Store,
	}

	client.interact()
}
