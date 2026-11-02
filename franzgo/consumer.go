package franzgo

import (
	"context"
	"fmt"
	"math/rand"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/twmb/franz-go/pkg/kgo"
)

func NewConsumer(brokers string) *kgo.Client {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(strings.Split(brokers, ",")...),
		kgo.ConsumerGroup(fmt.Sprintf("group-%d", rand.Intn(10000))),
		kgo.ConsumeTopics(viper.GetString("kafka.topic")),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()),
	)
	if err != nil {
		log.WithError(err).Panic("Unable to start the consumer")
	}
	return client
}

func PrepareConsume(client *kgo.Client, numMessages int) func() {
	log.Infof("Preparing to receive %d messages", numMessages)
	return func() {
		ctx := context.Background()
		consumed := 0
		for consumed < numMessages {
			fetches := client.PollFetches(ctx)
			if errs := fetches.Errors(); len(errs) > 0 {
				log.WithError(errs[0].Err).Panic("Unable to consume the message")
			}
			consumed += fetches.NumRecords()
		}
		log.Infof("Consumed %d messages successfully...", consumed)
	}
}
