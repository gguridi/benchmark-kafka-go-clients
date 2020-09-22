package sarama

import (
	"strings"
	"time"

	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Channel to detect when the client is do with the process.
	Done = make(chan bool)
)

// NewProducer returns a new Sarama async producer.
func NewProducer(brokers string) sarama.AsyncProducer {
	config := sarama.NewConfig()
	config.Version = sarama.V1_0_0_0
	config.Producer.Return.Successes = true
	config.Producer.Flush.Frequency = time.Duration(100) * time.Millisecond
	sarama.MaxRequestSize = 999000

	log.Debugf("Connecting to %s", brokers)
	producer, err := sarama.NewAsyncProducer(strings.Split(brokers, ","), config)
	if err != nil {
		log.WithError(err).Panic("Unable to start the producer")
	}
	return producer
}

// Prepare returns a function that can be used during the benchmark as it only
// performs the sending of messages, checking that the sending was successful.
func Prepare(producer sarama.AsyncProducer, message []byte, numMessages int) func() {
	log.Debugf("Preparing to send message of %d bytes %d times", len(message), numMessages)

	go func() {
		var nomessages int
		for _ = range producer.Successes() {
			nomessages++
			if nomessages%numMessages == 0 {
				log.Debugf("Sent %d messages... stopping...", nomessages)
				Done <- true
			}
		}
	}()

	go func() {
		for err := range producer.Errors() {
			log.WithError(err).Panic("Unable to deliver the message")
		}
	}()

	return func() {
		for j := 0; j < numMessages; j++ {
			producer.Input() <- &sarama.ProducerMessage{
				Topic:     viper.GetString("kafka.topic"),
				Partition: viper.GetInt32("kafka.partition"),
				Value:     sarama.ByteEncoder(message),
			}
		}
	}
}
