package confluent

import (
	"fmt"
	"math/rand"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewConsumer returns a new confluent consumer.
func NewConsumer(brokers string, poll bool) *kafka.Consumer {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               brokers,
		"group.id":                        fmt.Sprintf("group-%d", rand.Intn(10000)),
		"session.timeout.ms":              10000,
		"go.events.channel.enable":        !poll,
		"go.application.rebalance.enable": false,
		"enable.auto.commit":              false,
		"default.topic.config":            kafka.ConfigMap{"auto.offset.reset": "earliest"},
	})
	if err != nil {
		log.WithError(err).Panic("Unable to start the consumer")
	}

	topic := viper.GetString("kafka.topic")
	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.WithError(err).Panic("Unable to subscribe to the topic")
	}
	return consumer
}

// PreparePoll returns a function that can be used during the benchmark as it only
// performs the consuming of messages. It uses the poll/function method.
func PreparePoll(consumer *kafka.Consumer, numMessages int) func() {
	log.Infof("Preparing to receive %d messages", numMessages)
	return func() {
		var counter = 0
		for counter < numMessages {
			if ev := consumer.Poll(100); ev != nil {
				switch e := ev.(type) {
				case *kafka.Message:
					counter++
				case kafka.PartitionEOF:
					log.Panic("Reached Partition EOF: %v", ev)
				case kafka.Error:
					log.WithError(e).Panic("Unable to consume the message")
				default:
					log.Panic("Unable to consume the message: %v", ev)
				}
			}
		}
		log.Infof("Consumed %d messages successfully...", counter)
	}
}

// PrepareChannel returns a function that can be used during the benchmark as it only
// performs the consuming of messages. It uses the deprecated channel method.
func PrepareChannel(consumer *kafka.Consumer, numMessages int) func() {
	log.Infof("Preparing to receive %d messages", numMessages)
	return func() {
		var counter = 0
		for counter < numMessages {
			select {
			case ev := <-consumer.Events():
				switch ev.(type) {
				case *kafka.Message:
					counter++
				case kafka.PartitionEOF:
					log.Panic("Reached Partition EOF: %v", ev)
				case kafka.Error:
					log.WithError(ev.(kafka.Error)).Panic("Unable to consume the message")
				}
			}
		}
		log.Infof("Consumed %d messages successfully...", counter)
	}
}
