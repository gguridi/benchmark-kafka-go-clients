package kafkago

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	kafkago "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	BATCH_TIMEOUT = 100 * time.Millisecond
	BATCH_SIZE    = 10000
	BATCH_BYTES   = 999000
)

var (
	// Channel to detect when the client is done with the process.
	Done = make(chan bool)
)

// NewProducer returns a new kafkago asynchronous writer that signals Done once
// numMessages have been acknowledged by the brokers.
func NewProducer(brokers string, numMessages int) *kafkago.Writer {
	var delivered atomic.Int64
	return &kafkago.Writer{
		Addr:         kafkago.TCP(strings.Split(brokers, ",")...),
		Topic:        viper.GetString("kafka.topic"),
		Balancer:     &kafkago.RoundRobin{},
		BatchTimeout: BATCH_TIMEOUT,
		BatchSize:    BATCH_SIZE,
		BatchBytes:   BATCH_BYTES,
		RequiredAcks: kafkago.RequireAll,
		Async:        true,
		Completion: func(messages []kafkago.Message, err error) {
			if err != nil {
				log.WithError(err).Panic("Unable to deliver the message")
			}
			total := delivered.Add(int64(len(messages)))
			if total >= int64(numMessages) && total-int64(len(messages)) < int64(numMessages) {
				log.Infof("Sent %d messages... stopping...", total)
				Done <- true
			}
		},
	}
}

// Prepare returns a function that can be used during the benchmark as it only
// performs the sending of messages.
func Prepare(writer *kafkago.Writer, message []byte, numMessages int) func() {
	log.Infof("Preparing to send message of %d bytes %d times", len(message), numMessages)
	return func() {
		for range numMessages {
			err := writer.WriteMessages(context.Background(), kafkago.Message{Value: message})
			if err != nil {
				log.WithError(err).Panic("Unable to enqueue the message")
			}
		}
	}
}
