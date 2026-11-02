import { Empty } from "antd";
import { groupByMessageSize, LIBRARIES, selectBenchmarks } from "@/benchmarks";
import LibraryLegend from "@/components/LibraryLegend";
import ResultsTable from "@/components/ResultsTable";
import ThroughputCard from "@/components/ThroughputCard";
import { RESULTS } from "@/results";

export default function BenchmarkPage({ type, scheme }) {
    const benchmarks = selectBenchmarks({ results: RESULTS, type });
    if (benchmarks.length === 0) {
        return <Empty className="py-16" description="No benchmark results available yet" />;
    }
    const libraries = LIBRARIES[type];
    return (
        <div className="flex flex-col gap-6">
            <LibraryLegend libraries={libraries} scheme={scheme} />
            <div className="grid gap-6 lg:grid-cols-2">
                {groupByMessageSize(benchmarks).map((group) => (
                    <ThroughputCard
                        key={group.messageSize}
                        group={group}
                        libraries={libraries}
                        scheme={scheme}
                    />
                ))}
            </div>
            <ResultsTable benchmarks={benchmarks} scheme={scheme} />
        </div>
    );
}
