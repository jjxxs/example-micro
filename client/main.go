package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/etcdv3/v2"
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
			fmt.Println(err)
		} else {
			fmt.Println(rsp.GetGreeting())
		}

		fmt.Print("sleeping...")
		time.Sleep(2 * time.Second)
		fmt.Println("awake")
	}
}

func main() {
	registry := etcdv3.NewRegistry()

	service := micro.NewService(micro.Registry(registry))
	service.Init()

	// create the greeter client using the service name and client
	client := Client{
		greeter: api.NewGreeterService("greeter", service.Client()),
	}

	client.interact()
}
