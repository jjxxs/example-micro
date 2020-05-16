package client

import (
	"context"
	"strconv"
	"time"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	"github.com/vesose/example-micro/api"
)

type Client struct {
	greeter api.GreeterService
	store   store.Store
}

func New(greeter api.GreeterService, store store.Store) *Client {
	return &Client{
		greeter: greeter,
		store:   store,
	}
}

func (c *Client) Interact() {
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
