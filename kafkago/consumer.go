package kafkago

import (
	"context"
	"fmt"
	"math/rand"

	kafkago "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewReader returns a new kafkago writer.
func NewReader(brokers string) *kafkago.Reader {
	return kafkago.NewReader(kafkago.ReaderConfig{
		Brokers:  []string{brokers},
		GroupID:  fmt.Sprintf("group-%d", rand.Intn(10000)),
		Topic:    viper.GetString("kafka.topic"),
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

// PrepareReader returns a function that can be used during the benchmark as it only
// performs the consuming of messages.
func PrepareReader(reader *kafkago.Reader, numMessages int) func() {
	log.Debugf("Preparing to receive %d messages", numMessages)
	return func() {
		ctx := context.Background()
		count := 0
		for {
			if _, err := reader.ReadMessage(ctx); err != nil {
				log.WithError(err).Panic("Unable to consume the message")
			}
			if count >= numMessages {
				log.Infof("Consumed %d messages successfully...", numMessages)
				return
			}
			count++
		}
	}
}
