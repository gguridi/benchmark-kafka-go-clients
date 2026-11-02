import LibrarySwatch from "@/components/LibrarySwatch";
import { formatMegabytes, formatMessageCount, formatSeconds, formatThroughput } from "@/format";
import { libraryLabel } from "@/theme";

export default function ThroughputTooltip({ active, payload, scheme }) {
    if (!active || !payload?.length) {
        return null;
    }
    const row = payload[0].payload;
    const benchmarks = Object.values(row)
        .filter((entry) => typeof entry === "object")
        .sort((first, second) => second.throughput - first.throughput);
    return (
        <div className="flex flex-col gap-2 rounded-md border border-hairline bg-surface p-3 text-xs shadow-sm">
            <p className="font-semibold text-ink">{formatMessageCount(row.numMessages)}</p>
            <table className="tabular-nums">
                <tbody>
                    {benchmarks.map((benchmark) => (
                        <tr key={benchmark.library}>
                            <td className="pr-3">
                                <span className="flex items-center gap-2 text-ink">
                                    <LibrarySwatch library={benchmark.library} scheme={scheme} />
                                    {libraryLabel(benchmark.library)}
                                </span>
                            </td>
                            <td className="pr-3 text-right text-ink">
                                {formatThroughput(benchmark.throughput)}
                            </td>
                            <td className="pr-3 text-right text-ink-secondary">
                                {formatMegabytes(benchmark.megabytesPerSecond)}
                            </td>
                            <td className="text-right text-ink-secondary">
                                {formatSeconds(benchmark.averageSeconds)}
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}
