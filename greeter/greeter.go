package greeter

import (
	"context"

	"github.com/micro/go-micro/v2/logger"
	"github.com/vesose/example-micro/api"
)

type Greeter struct {
	counter api.HelloCounterService
}

func New(counter api.HelloCounterService) *Greeter {
	return &Greeter{
		counter: counter,
	}
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
