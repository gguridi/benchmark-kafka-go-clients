package franzgo

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/twmb/franz-go/pkg/kgo"
)

const LINGER = 100 * time.Millisecond

var Done = make(chan bool)

func NewProducer(brokers string) *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(strings.Split(brokers, ",")...),
		kgo.DefaultProduceTopic(viper.GetString("kafka.topic")),
		kgo.ProducerLinger(LINGER),
	)
	if err != nil {
		log.WithError(err).Panic("Unable to start the producer")
	}
	return client
}

func Prepare(client *kgo.Client, message []byte, numMessages int) func() {
	log.Infof("Preparing to send message of %d bytes %d times", len(message), numMessages)
	var delivered atomic.Int64
	onDelivery := func(_ *kgo.Record, err error) {
		if err != nil {
			log.WithError(err).Panic("Unable to deliver the message")
		}
		if delivered.Add(1) == int64(numMessages) {
			log.Infof("Sent %d messages... stopping...", numMessages)
			Done <- true
		}
	}
	return func() {
		ctx := context.Background()
		for range numMessages {
			client.Produce(ctx, &kgo.Record{Value: message}, onDelivery)
		}
	}
}
