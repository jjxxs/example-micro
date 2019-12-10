package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	api "github.com/vesose/example-micro/api"
)

type Greeter struct {
	counter api.HelloCounterService
}

func (g *Greeter) Hello(ctx context.Context, req *api.HelloRequest, rsp *api.HelloResponse) error {

	counterRsp, err := g.counter.Inc(context.TODO(), &api.IncRequest{
		Name: req.Name,
	})

	if err != nil {
		fmt.Println(err)
		rsp.Greeting = "Hiho " + req.Name
	} else {
		fmt.Printf("Counter = %d\n", counterRsp.Counter)
		if counterRsp.GetCounter()%2 == 1 {
			rsp.Greeting = "Hello " + req.Name
		} else {
			rsp.Greeting = "See you " + req.Name
		}
	}

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
	)

	service.Init()

	counter := micro.NewService()
	counter.Init()

	api.RegisterGreeterHandler(service.Server(), &Greeter{
		counter: api.NewHelloCounterService("counter", counter.Client()),
	})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
