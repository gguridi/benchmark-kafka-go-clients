package main

import (
	"flag"
	"fmt"

	"github.com/gguridi/benchmark-kafka-go-clients/confluent"
	"github.com/gguridi/benchmark-kafka-go-clients/sarama"
	. "github.com/onsi/ginkgo"
	log "github.com/sirupsen/logrus"
)

var _ = Describe("Benchmarks", func() {

	Measure("producer", func(b Benchmarker) {
		flag.Parse()
		name := fmt.Sprintf("%s processing %d messages of %d bytes size", Library, NumMessages, MessageSize)
		switch Library {
		case "confluent":
			process, finish := GetConfluent()
			b.Time(name, func() {
				process()
				<-confluent.Done
			})
			finish()
			break
		case "sarama":
			process, finish := GetSarama()
			b.Time(name, func() {
				process()
				<-sarama.Done
			})
			if err := finish(); err != nil {
				log.WithError(err).Panic("Unable to close the producer")
			}
			break
		case "kafkago":
			if NumMessages <= 1000 {
				process, finish := GetKafkaGo()
				b.Time(name, func() {
					process()
				})
				if err := finish(); err != nil {
					log.WithError(err).Panic("Unable to close the producer")
				}
			}
			break
		default:
			log.Panicf("Unable to find the libray %+v", Library)
		}
	}, 10)
})
