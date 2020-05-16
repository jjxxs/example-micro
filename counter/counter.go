package counter

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	"github.com/vesose/example-micro/api"
)

type Counter struct {
	publisher micro.Event
	counters  map[string]int32
}

func New(publisher micro.Event) *Counter {
	return &Counter{
		publisher: publisher,
		counters:  make(map[string]int32),
	}
}

func (g *Counter) Inc(ctx context.Context, req *api.IncRequest, rsp *api.SumResponse) error {
	name := req.Name
	g.counters[name]++
	rsp.Counter = g.counters[name]
	msg := fmt.Sprintf("Request for %s, new counter = %d", name, g.counters[name])

	logger.Info(msg)

	uuid, err := uuid.NewRandom()

	if err != nil {
		logger.Errorf("error creating uuid: %v", err)
	}

	if err := g.publisher.Publish(context.Background(), &api.Event{
		Id:        uuid.String(),
		Timestamp: time.Now().Unix(),
		Message:   msg,
	}); err != nil {
		logger.Errorf("error while publishing: %v", err)
	}

	return nil
}
