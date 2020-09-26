# Benchmark of Kafka Go Clients

![Benchmark](https://github.com/gguridi/benchmark-kafka-go-clients/workflows/Benchmark/badge.svg?branch=master)

Benchmark to test different kafka go clients to compare them under the same conditions.

This started as a proof of concept of creating some benchmarks that could run directly
on github actions, publishing the results in a dashboard associated with the same repository.

This benchmark has been greatly inspired by Matt Howlett's [benchmark](https://gist.github.com/mhowlett/e9491aad29817aeda6003c3404874b35). I shamelessly used what it has instead of reading the documentation of the different clients :S.

## Clients

The clients tested in this benchmark are:

- [Confluent-kafka-Go](github.com/confluentinc/confluent-kafka-go)
- [Sarama](github.com/Shopify/sarama)
- [KafkaGo](github.com/segmentio/kafka-go)

## Setup

In order to run the benchmarks we need the following cli tools installed:

- [Ginkgo](https://github.com/onsi/ginkgo).

## Running

In order to run the benchmark tests we need a working kafka image. In this benchmark I've used
docker-based Kafka instances but it's possible to use any kafka.

To run the benchmarks we use the following syntax:

```bash
ginkgo -focus=${type} ./... -- -test.bench=. -library=${client} -num=${num-messages} -size=${size-messages} -brokers=${brokers} -topic=${topic}
```

### Type

It can be one of the following:

- *producer*: To test the client producers against the kafka brokers selected.
- *consumer*: To test the client consumers against the kafka brokers selected.
- *lag*: To perform a lag check agains a topic to see if the producer sent all the messages as expected. 

The syntax of the lag is slightly different to the other two types.

```bash
ginkgo -focus=lag ./... -- -test.bench=. -num=${num-messages} -brokers=${brokers} -topic=${topic}
```

Where `num-messages` is the number of messages we would expect to find in that topic.

### Client

It can be one of the following:

- *confluent*: (producer only). To test the Confluent-kafka-Go producer.
- *confluent-poll*: (consumer only). To test the Confluent-kafka-Go polling-based consumer.
- *confluent-channel*: (consumer only). To test the Confluent-kafka-Go channel-based (and deprecated) consumer.
- *sarama*: To test the Sarama producer/consumer.
- *kafkago*: To test the KafkaGo producer/consumer.

### Num-Messages

It's an integer specifying the number of messages to produce/consume for the benchmark.

### Size-Messages

It's an integer specifying the size of messages to use for the benchmark. In case of the
consumers the Kafka instance will be prepopulated with messages of this size to be consumed.

### Brokers

Location of the kafka brokers to use for this benchmark. It includes the port. For instance `localhost:9092`.

### Topic

The topic as string to use for this benchmark.

## Output

The output of the benchmarks is automatically display in the dashboard. The results
are downloaded from the github actions, where are stored as artifacts, and used
as sources when building the results dashboard.

- https://gguridi.github.io/benchmark-kafka-go-clients/

If a client doesn't appear in the graph is because the time used in the benchmark is so much
higher than I didn't consider it necessary to perform further tests of it.

## Thanks

Special thanks to the components that I used for this:

- Matt Howlett's [Benchmark](https://gist.github.com/mhowlett/e9491aad29817aeda6003c3404874b35)
- Creative Tim [Vue Material Dashboard](https://www.creative-tim.com/product/vue-material-dashboard)
