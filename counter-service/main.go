package main

import (
	"context"
	"fmt"
	"log"

	"github.com/micro/go-micro"
	api "github.com/vesose/example-micro/api"
)

type Counter struct {
	counters map[string]int32
}

func (g *Counter) Inc(ctx context.Context, req *api.IncRequest, rsp *api.SumResponse) error {
	name := req.Name
	g.counters[name]++
	rsp.Counter = g.counters[name]
	fmt.Printf("Request for %s, new counter = %d\n", name, g.counters[name])
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("counter"),
		micro.Version("latest"),
	)

	service.Init()

	api.RegisterHelloCounterHandler(service.Server(), &Counter{counters: make(map[string]int32)})

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
