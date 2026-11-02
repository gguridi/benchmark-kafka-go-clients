package kafkago

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	kafkago "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const COMMIT_INTERVAL = time.Second

// NewReader returns a new kafkago reader.
func NewReader(brokers string) *kafkago.Reader {
	return kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:        strings.Split(brokers, ","),
		GroupID:        fmt.Sprintf("group-%d", rand.Intn(10000)),
		Topic:          viper.GetString("kafka.topic"),
		StartOffset:    kafkago.FirstOffset,
		CommitInterval: COMMIT_INTERVAL,
		MinBytes:       10e3, // 10KB
		MaxBytes:       10e6, // 10MB
	})
}

// PrepareReader returns a function that can be used during the benchmark as it only
// performs the consuming of messages.
func PrepareReader(reader *kafkago.Reader, numMessages int) func() {
	log.Infof("Preparing to receive %d messages", numMessages)
	return func() {
		ctx := context.Background()
		for range numMessages {
			if _, err := reader.ReadMessage(ctx); err != nil {
				log.WithError(err).Panic("Unable to consume the message")
			}
		}
		log.Infof("Consumed %d messages successfully...", numMessages)
	}
}
