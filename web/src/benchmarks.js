export const PRODUCER = "producer";
export const CONSUMER = "consumer";

export const LIBRARIES = {
    [PRODUCER]: ["confluent", "sarama", "kafkago", "franzgo"],
    [CONSUMER]: ["confluent-poll", "sarama", "kafkago", "franzgo"],
};

const BYTES_PER_MEGABYTE = 1_000_000;

export function selectBenchmarks({ results, type }) {
    return results
        .filter((result) => result.type === type && result.average > 0)
        .map(toBenchmark)
        .sort(byConfigurationThenThroughput);
}

export function groupByMessageSize(benchmarks) {
    const sizes = [...new Set(benchmarks.map(({ messageSize }) => messageSize))];
    return sizes.map((messageSize) => {
        const ofSize = benchmarks.filter((benchmark) => benchmark.messageSize === messageSize);
        return { messageSize, rows: toChartRows(ofSize), fastest: fastestPerMessageCount(ofSize) };
    });
}

function toBenchmark(result) {
    const throughput = result.numMessages / result.average;
    return {
        library: result.library,
        numMessages: result.numMessages,
        messageSize: result.sizeMessages,
        averageSeconds: result.average,
        stdDeviationSeconds: result.stdDeviation,
        samples: result.numberOfSamples,
        throughput,
        megabytesPerSecond: (throughput * result.sizeMessages) / BYTES_PER_MEGABYTE,
    };
}

function byConfigurationThenThroughput(first, second) {
    return (
        first.messageSize - second.messageSize ||
        first.numMessages - second.numMessages ||
        second.throughput - first.throughput
    );
}

function toChartRows(benchmarks) {
    const counts = [...new Set(benchmarks.map(({ numMessages }) => numMessages))];
    return counts.map((numMessages) => {
        const ofCount = benchmarks.filter((benchmark) => benchmark.numMessages === numMessages);
        return Object.fromEntries([
            ["numMessages", numMessages],
            ...ofCount.map((benchmark) => [benchmark.library, benchmark]),
        ]);
    });
}

function fastestPerMessageCount(benchmarks) {
    const counts = [...new Set(benchmarks.map(({ numMessages }) => numMessages))];
    return counts.map((numMessages) =>
        benchmarks.find((benchmark) => benchmark.numMessages === numMessages),
    );
}
