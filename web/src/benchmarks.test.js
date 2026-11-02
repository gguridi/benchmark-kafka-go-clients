import { describe, expect, it } from "vitest";
import { CONSUMER, groupByMessageSize, PRODUCER, selectBenchmarks } from "./benchmarks";

const result = (overrides) => ({
    type: PRODUCER,
    library: "sarama",
    numMessages: 10000,
    sizeMessages: 1000,
    numberOfSamples: 3,
    average: 2,
    stdDeviation: 0.1,
    ...overrides,
});

const SARAMA_SMALL = result({});
const FRANZGO_SMALL = result({ library: "franzgo", average: 1 });
const SARAMA_LARGE = result({ numMessages: 1000000, average: 50 });
const SARAMA_BIG_MESSAGES = result({ sizeMessages: 5000, average: 4 });
const CONSUMER_RESULT = result({ type: CONSUMER });
const EMPTY_RESULT = result({ library: "kafkago", average: 0 });

describe("selectBenchmarks", () => {
    const benchmarks = selectBenchmarks({
        results: [SARAMA_LARGE, SARAMA_SMALL, CONSUMER_RESULT, EMPTY_RESULT, FRANZGO_SMALL],
        type: PRODUCER,
    });

    it("keeps only results of the requested type with a measured time, fastest first", () => {
        expect(
            benchmarks.map(({ library, numMessages }) => `${library}-${numMessages}`),
        ).toStrictEqual(["franzgo-10000", "sarama-10000", "sarama-1000000"]);
    });

    it("computes the throughput in messages per second", () => {
        expect(benchmarks[1].throughput).toBe(5000);
    });

    it("computes the throughput in megabytes per second", () => {
        expect(benchmarks[1].megabytesPerSecond).toBe(5);
    });
});

describe("groupByMessageSize", () => {
    const groups = groupByMessageSize(
        selectBenchmarks({
            results: [SARAMA_SMALL, FRANZGO_SMALL, SARAMA_LARGE, SARAMA_BIG_MESSAGES],
            type: PRODUCER,
        }),
    );

    it("creates one group per message size", () => {
        expect(groups.map(({ messageSize }) => messageSize)).toStrictEqual([1000, 5000]);
    });

    it("creates one chart row per message count keyed by library", () => {
        expect(
            groups[0].rows.map((row) => [row.numMessages, Object.keys(row).sort()]),
        ).toStrictEqual([
            [10000, ["franzgo", "numMessages", "sarama"]],
            [1000000, ["numMessages", "sarama"]],
        ]);
    });

    it("picks the fastest library for each message count", () => {
        expect(groups[0].fastest.map(({ library }) => library)).toStrictEqual([
            "franzgo",
            "sarama",
        ]);
    });
});
