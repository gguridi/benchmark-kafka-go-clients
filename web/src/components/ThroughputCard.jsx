import ThroughputChart from "@/components/ThroughputChart";
import { formatMessageCount, formatMessageSize, formatThroughput } from "@/format";
import { libraryLabel } from "@/theme";

export default function ThroughputCard({ group, libraries, scheme }) {
    return (
        <section className="flex flex-col gap-3 rounded-lg border border-hairline bg-surface p-5">
            <div className="flex flex-col gap-1">
                <h2 className="text-base font-semibold">
                    {formatMessageSize(group.messageSize)} messages
                </h2>
                <ul className="text-sm text-ink-secondary">
                    {group.fastest.map((benchmark) => (
                        <li key={benchmark.numMessages}>
                            Fastest at {formatMessageCount(benchmark.numMessages)}:{" "}
                            <span className="font-medium text-ink">
                                {libraryLabel(benchmark.library)}
                            </span>{" "}
                            ({formatThroughput(benchmark.throughput)})
                        </li>
                    ))}
                </ul>
            </div>
            <ThroughputChart group={group} libraries={libraries} scheme={scheme} />
        </section>
    );
}
