#!/usr/bin/env bash
set -uo pipefail

if [[ $# -ne 3 ]]; then
    echo "Usage: $0 <producer|consumer> <num-messages> <message-size>" >&2
    exit 2
fi

TYPE=$1
NUM_MESSAGES=$2
MESSAGE_SIZE=$3
BROKERS=${BROKERS:-localhost:9092}
KAFKA_CONTAINER=${KAFKA_CONTAINER:-benchmark-kafka}
OUTPUT_DIR=${OUTPUT_DIR:-results}
GINKGO=${GINKGO:-ginkgo}
PRODUCERS=(sarama confluent kafkago franzgo)
CONSUMERS=(sarama confluent-poll kafkago franzgo)
FAILED=()

export ACK_GINKGO_DEPRECATIONS=1.16.5

kafka_topics() {
    docker exec "$KAFKA_CONTAINER" /opt/kafka/bin/kafka-topics.sh --zookeeper localhost:2181 "$@"
}

create_topic() {
    for _ in $(seq 1 60); do
        kafka_topics --create --if-not-exists --partitions 8 --replication-factor 1 --topic "$1" >/dev/null 2>&1
        has_all_leaders "$1" && return 0
        sleep 2
    done
    echo "Unable to create topic $1 with a leader on every partition" >&2
    exit 1
}

has_all_leaders() {
    [[ $(kafka_topics --describe --topic "$1" 2>/dev/null | grep -c "Leader: [0-9]") -eq 8 ]]
}

run_spec() {
    local focus=$1 library=$2 topic=$3
    "$GINKGO" -focus="$focus" . -- -library="$library" -num="$NUM_MESSAGES" -size="$MESSAGE_SIZE" -brokers="$BROKERS" -topic="$topic"
}

run_library() {
    local library=$1 topic=$2
    echo "::group::$TYPE $library ($NUM_MESSAGES messages of $MESSAGE_SIZE bytes)"
    rm -f results.json
    if run_spec "$TYPE" "$library" "$topic" && [[ -s results.json ]]; then
        mv results.json "$OUTPUT_DIR/$TYPE-$library-$NUM_MESSAGES-$MESSAGE_SIZE.json"
    else
        FAILED+=("$library")
    fi
    echo "::endgroup::"
}

run_producers() {
    for library in "${PRODUCERS[@]}"; do
        local topic="benchmark-$library"
        create_topic "$topic"
        run_library "$library" "$topic"
        kafka_topics --delete --topic "$topic"
    done
}

run_consumers() {
    local topic="benchmark-consumers"
    create_topic "$topic"
    run_spec "prepopulates" "sarama" "$topic" || exit 1
    for library in "${CONSUMERS[@]}"; do
        run_library "$library" "$topic"
    done
}

mkdir -p "$OUTPUT_DIR"
case "$TYPE" in
    producer) run_producers ;;
    consumer) run_consumers ;;
    *) echo "Unknown benchmark type: $TYPE" >&2; exit 2 ;;
esac

if [[ ${#FAILED[@]} -gt 0 ]]; then
    echo "Failed libraries: ${FAILED[*]}" >&2
    exit 1
fi
