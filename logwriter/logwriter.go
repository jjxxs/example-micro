package logwriter

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/store"
	"github.com/vesose/example-micro/api"
)

type LogWriter struct {
	store store.Store
}

func New(store store.Store) *LogWriter {
	return &LogWriter{store: store}
}

func (s *LogWriter) Process(ctx context.Context, event *api.Event) error {
	logger.Infof("Received event msg: %+v", event.GetMessage())

	newSleep := []byte(fmt.Sprintf("%v", rand.Intn(2000))) // nolint:gomnd

	record := store.Record{
		Key:    "sleep",
		Value:  newSleep,
		Expiry: 0,
	}
	err := s.store.Write(&record, func(o *store.WriteOptions) { o.Table = "sleeper" })

	if err != nil {
		logger.Infof("error while writing to store: %+v", err)
	}

	return err
}
