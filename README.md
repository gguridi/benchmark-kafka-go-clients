# Benchmark of Kafka Go Clients

![Benchmark](https://github.com/gguridi/benchmark-kafka-go-clients/workflows/Benchmark/badge.svg?branch=master)

Benchmark to test different kafka go clients to compare them under the same conditions.

This started as a proof of concept of creating some benchmarks that could run directly
on github actions, publishing the results in a dashboard associated with the same repository.

This benchmark has been greatly inspired by Matt Howlett's [benchmark](https://gist.github.com/mhowlett/e9491aad29817aeda6003c3404874b35). I shamelessly used what it has instead of reading the documentation of the different clients :S.

## Clients

The clients tested in this benchmark are:

- [Confluent-kafka-go](github.com/confluentinc/confluent-kafka-go)
- [Sarama](github.com/Shopify/sarama)
- [KafkaGo](github.com/segmentio/kafka-go)

## Output

The output of the benchmarks is automatically display in the dashboard. The results
are downloaded from the github actions, where are stored as artifacts, and used
as sources when building the results dashboard.

- https://gguridi.github.io/benchmark-kafka-go-clients/

## Thanks

Special thanks to the components that I used for this:

- Matt Howlett's [Benchmark](https://gist.github.com/mhowlett/e9491aad29817aeda6003c3404874b35)
- Creative Tim [Vue Material Dashboard](https://www.creative-tim.com/product/vue-material-dashboard)
