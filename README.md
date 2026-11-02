# Benchmark of Kafka Go Clients

![Benchmark](https://github.com/gguridi/benchmark-kafka-go-clients/workflows/Benchmark/badge.svg?branch=master)

Throughput benchmarks of the Kafka Go clients, run on GitHub Actions and
published to a [dashboard](https://gguridi.github.io/benchmark-kafka-go-clients/).

Inspired by Matt Howlett's
[benchmark](https://gist.github.com/mhowlett/e9491aad29817aeda6003c3404874b35).

## Clients

| Client                                                                   | Producer    | Consumer         |
| ------------------------------------------------------------------------ | ----------- | ---------------- |
| [confluent-kafka-go](https://github.com/confluentinc/confluent-kafka-go) | `confluent` | `confluent-poll` |
| [Sarama](https://github.com/IBM/sarama)                                  | `sarama`    | `sarama`         |
| [kafka-go](https://github.com/segmentio/kafka-go)                        | `kafkago`   | `kafkago`        |
| [franz-go](https://github.com/twmb/franz-go)                             | `franzgo`   | `franzgo`        |

## Prerequisites

- Go (version in `go.mod`)
- Ginkgo v1 CLI
- Docker
- Node 24 and Yarn 1, for the dashboard

## Run the benchmarks

Start Kafka:

```bash
docker run --rm --name benchmark-kafka -e KAFKA__ADVERTISED_HOST_NAME=localhost -d -p 9092:9092 -p 2181:2181 gguridi/kafka:2.12-2.0.0
```

Run every client for one configuration. Results are written to
`results/<type>-<client>-<num>-<size>.json`:

```bash
scripts/benchmark.sh <producer|consumer> <num-messages> <message-size>
```

The script reads `BROKERS` (default `localhost:9092`), `KAFKA_CONTAINER` (default
`benchmark-kafka`), `OUTPUT_DIR` (default `results`) and `GINKGO` (default `ginkgo`).

Run a single client:

```bash
ginkgo -focus=producer . -- -library=franzgo -num=10000 -size=1000 -brokers=localhost:9092 -topic=test-topic
```

Consumers read messages already in the topic; fill it first with
`-focus=prepopulates`, which writes 1.5 × `-num` messages.

## Dashboard

```bash
cd web
yarn install
yarn dev
```

The dashboard reads the result files from `web/src/results/`.

```bash
yarn lint
yarn test
yarn build
```
