package main

import (
	"flag"
	"math/rand"
	"time"

	log "github.com/sirupsen/logrus"
)

var (
	// Library contains the library name we are going to run the benchmarks for.
	Library string
	// Brokers contains the list of brokers, comma-separated, to use.
	Brokers string
	// Topic contains the topic to use in this test.
	Topic string
	// NumMessages contains the number of messages to send.
	NumMessages int
	// MessageSize contains the size of the message to send.
	MessageSize int
	// Rune to use in the random string generator
	Characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func init() {
	rand.Seed(time.Now().UnixNano())
	flag.StringVar(&Brokers, "brokers", "", "Brokers to use for this benchmark")
	flag.StringVar(&Topic, "topic", "", "Topic to use for this benchmark")
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

func main() {
	log.Panic("This benchmark is intended to be run with ginkgo through tests")
}
