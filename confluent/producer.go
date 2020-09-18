package confluent

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Channel to detect when the client is do with the process.
	Done = make(chan bool)
)

func NewProducer(brokers string) *kafka.Producer {
	config := &kafka.ConfigMap{"bootstrap.servers": brokers, "linger.ms": 100}
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.WithError(err).Panic("Unable to start the producer")
	}
	return producer
}

func Prepare(producer *kafka.Producer, message []byte, numMessages int) func() {
	topic := viper.GetString("kafka.topic")
	log.Debugf("Preparing to send message of %d bytes %d times", len(message), numMessages)

	go func() {
		var msgCount int
		for e := range producer.Events() {
			msg := e.(*kafka.Message)
			if msg.TopicPartition.Error != nil {
				log.WithError(msg.TopicPartition.Error).Panic("Unable to deliver the message")
			}
			msgCount++
			if msgCount >= numMessages {
				log.Debugf("Sent %d messages... stopping...", msgCount)
				Done <- true
			}
		}
	}()

	return func() {
		for j := 0; j < numMessages; j++ {
			producer.ProduceChannel() <- &kafka.Message{
				TopicPartition: kafka.TopicPartition{
					Topic:     &topic,
					Partition: viper.GetInt32("kafka.partition"),
				},
				Value: message,
			}
		}
	}
}
