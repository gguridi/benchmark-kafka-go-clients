package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/gguridi/benchmark-kafka-go-clients/confluent"
	"github.com/gguridi/benchmark-kafka-go-clients/kafkago"
	"github.com/gguridi/benchmark-kafka-go-clients/sarama"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Library contains the library name we are going to run the benchmarks for.
	Library string
	// NumMessages contains the number of messages to send.
	NumMessages int
	// MessageSize contains the size of the message to send.
	MessageSize int
	// Rune to use in the random string generator
	Characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.StringVar(&Library, "library", "", "Library to use for this benchmark")
	flag.IntVar(&NumMessages, "num", 1000, "Number of messages to send")
	flag.IntVar(&MessageSize, "size", 1000, "Number of messages to send")
}

// GenMessage generates a random message of X bytes to use in the benchmarks.
func GenMessage() []byte {
	b := make([]rune, MessageSize)
	for i := range b {
		b[i] = Characters[rand.Intn(len(Characters))]
	}
	return []byte(string(b))
}

func GetConfluent() (func(), func()) {
	producer := confluent.NewProducer(viper.GetString("kafka.brokers"))
	return confluent.Prepare(producer, GenMessage(), NumMessages), producer.Close
}

func GetSarama() (func(), func() error) {
	producer := sarama.NewProducer(viper.GetString("kafka.brokers"))
	return sarama.Prepare(producer, GenMessage(), NumMessages), producer.Close
}

func GetKafkaGo() (func(), func() error) {
	writer := kafkago.NewProducer(viper.GetString("kafka.brokers"))
	return kafkago.Prepare(writer, GenMessage(), NumMessages), writer.Close
}

func main() {
	log.Panic("This benchmark is intended to be run with ginkgo through tests")
}
