package main

import (
	"context"
	"fmt"
	"time"

	"github.com/micro/go-micro"
	api "github.com/vesose/example-micro/api"
)

type Client struct {
	greeter api.GreeterService
}

func (c Client) interact() {
	for {
		rsp, err := c.greeter.Hello(context.TODO(), &api.HelloRequest{
			Name: "John",
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rsp.GetGreeting())
		}

		fmt.Print("sleeping...")
		time.Sleep(10 * time.Second)
		fmt.Println("awake")
	}
}

func main() {

	service := micro.NewService()
	service.Init()

	// create the greeter client using the service name and client
	client := Client{
		greeter: api.NewGreeterService("greeter", service.Client()),
	}

	client.interact()

}
