package main

import (
	"context"
	"log"

	"github.com/micro/go-micro"
	api "github.com/vesose/example-micro/api"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *api.HelloRequest, rsp *api.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	service.Init()

	api.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
