package confluent

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	log "github.com/sirupsen/logrus"
)

// Lag returns the LAG of certain consumer. If the consumer hasn't consumed any message
// and the consumer group id is new, the LAG would identify the number of messages
// available in that topic.
func Lag(consumer *kafka.Consumer) int {
	var lag int

	// Get the current assigned partitions.
	topicPartitions, err := consumer.Assignment()
	log.Infof("Found topic partitions: %v", topicPartitions)
	if err != nil {
		log.WithError(err).Panic("Unable to retrieve the topic partitions")
	}

	// Get the current offset for each partition, assigned to this consumer group.
	topicPartitions, err = consumer.Committed(topicPartitions, 5000)
	if err != nil {
		log.WithError(err).Panic("Unable to retrieve the committed offsets")
	}

	// Loop over the topic partitions, get the high watermark for each toppar, and
	// subtract the current offset from that number, to get the total "lag". We
	// combine this value for each toppar to get the final backlog integer.
	var low, high int64
	for i := range topicPartitions {
		topic := topicPartitions[i].Topic
		partition := topicPartitions[i].Partition
		low, high, err = consumer.QueryWatermarkOffsets(*topic, partition, 5000)
		log.Infof("Checked lag for %s/%d: result %d, %d", *topic, partition, low, high)
		if err != nil {
			log.WithError(err).Panic("Unable to retrieve the offsets for the partition %v", partition)
		}

		offset := int64(topicPartitions[i].Offset)
		if topicPartitions[i].Offset == kafka.OffsetInvalid {
			offset = low
		}

		lag = lag + int(high-offset)
	}
	return lag
}
