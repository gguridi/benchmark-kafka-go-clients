package main

import (
	"flag"
	"fmt"

	"github.com/gguridi/benchmark-kafka-go-clients/confluent"
	"github.com/gguridi/benchmark-kafka-go-clients/franzgo"
	"github.com/gguridi/benchmark-kafka-go-clients/kafkago"
	"github.com/gguridi/benchmark-kafka-go-clients/sarama"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const PREPOPULATE_FACTOR = 1.5

var _ = Describe("Benchmarks", func() {

	It("checks lag", func() {
		flag.Parse()
		consumer := confluent.NewConsumer(viper.GetString("kafka.brokers"))
		confluent.PreparePoll(consumer, 1)()
		Expect(confluent.Lag(consumer)).To(Equal(NumMessages))
	})

	It("prepopulates the topic", func() {
		flag.Parse()
		total := int(float64(NumMessages) * PREPOPULATE_FACTOR)
		log.Infof("Prepopulating Kafka with %d messages of %d bytes", total, MessageSize)
		producer := sarama.NewProducer(viper.GetString("kafka.brokers"))
		sarama.Prepare(producer, GenMessage(), total)()
		<-sarama.Done
		Expect(producer.Close()).To(Succeed())
	})

	Measure("producer", func(b Benchmarker) {
		flag.Parse()
		name := fmt.Sprintf("%s producing %d messages of %d bytes size", Library, NumMessages, MessageSize)
		switch Library {
		case "confluent":
			producer := confluent.NewProducer(viper.GetString("kafka.brokers"))
			process := confluent.Prepare(producer, GenMessage(), NumMessages)
			b.Time(name, func() {
				process()
				<-confluent.Done
			})
			producer.Close()
		case "sarama":
			producer := sarama.NewProducer(viper.GetString("kafka.brokers"))
			process := sarama.Prepare(producer, GenMessage(), NumMessages)
			b.Time(name, func() {
				process()
				<-sarama.Done
			})
			if err := producer.Close(); err != nil {
				log.WithError(err).Panic("Unable to close the producer")
			}
		case "kafkago":
			writer := kafkago.NewProducer(viper.GetString("kafka.brokers"), NumMessages)
			process := kafkago.Prepare(writer, GenMessage(), NumMessages)
			b.Time(name, func() {
				process()
				<-kafkago.Done
			})
			if err := writer.Close(); err != nil {
				log.WithError(err).Panic("Unable to close the producer")
			}
		case "franzgo":
			client := franzgo.NewProducer(viper.GetString("kafka.brokers"))
			process := franzgo.Prepare(client, GenMessage(), NumMessages)
			b.Time(name, func() {
				process()
				<-franzgo.Done
			})
			client.Close()
		default:
			log.Panicf("Unable to find the libray %+v", Library)
		}
	}, 3)

	Measure("consumer", func(b Benchmarker) {
		flag.Parse()
		name := fmt.Sprintf("%s consuming %d messages", Library, NumMessages)
		switch Library {
		case "confluent-poll":
			consumer := confluent.NewConsumer(viper.GetString("kafka.brokers"))
			process := confluent.PreparePoll(consumer, NumMessages)
			b.Time(name, process)
			consumer.Close()
		case "sarama":
			consumer := sarama.NewConsumer(viper.GetString("kafka.brokers"), NumMessages)
			process := sarama.PrepareConsume(consumer)
			b.Time(name, process)
			if err := consumer.Client.Close(); err != nil {
				log.WithError(err).Panic("Unable to close the consumer")
			}
		case "kafkago":
			reader := kafkago.NewReader(viper.GetString("kafka.brokers"))
			process := kafkago.PrepareReader(reader, NumMessages)
			b.Time(name, process)
			if err := reader.Close(); err != nil {
				log.WithError(err).Panic("Unable to close the consumer")
			}
		case "franzgo":
			client := franzgo.NewConsumer(viper.GetString("kafka.brokers"))
			process := franzgo.PrepareConsume(client, NumMessages)
			b.Time(name, process)
			client.Close()
		default:
			log.Panicf("Unable to find the libray %+v", Library)
		}
	}, 1)
})
