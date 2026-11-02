import { Table } from "antd";
import LibrarySwatch from "@/components/LibrarySwatch";
import {
    formatMegabytes,
    formatMessageCount,
    formatMessageSize,
    formatSeconds,
    formatThroughput,
} from "@/format";
import { libraryLabel } from "@/theme";

export default function ResultsTable({ benchmarks, scheme }) {
    return (
        <section className="flex flex-col gap-3">
            <h2 className="text-base font-semibold">All results</h2>
            <Table
                className="tabular-nums"
                size="small"
                pagination={false}
                scroll={{ x: true }}
                rowKey={({ library, numMessages, messageSize }) =>
                    `${library}-${numMessages}-${messageSize}`
                }
                dataSource={benchmarks}
                columns={columns(scheme)}
            />
        </section>
    );
}

function columns(scheme) {
    return [
        {
            title: "Client",
            dataIndex: "library",
            render: (library) => (
                <span className="flex items-center gap-2 whitespace-nowrap">
                    <LibrarySwatch library={library} scheme={scheme} />
                    {libraryLabel(library)}
                </span>
            ),
        },
        { title: "Message size", dataIndex: "messageSize", render: formatMessageSize },
        { title: "Messages", dataIndex: "numMessages", render: formatMessageCount },
        { title: "Throughput", dataIndex: "throughput", align: "right", render: formatThroughput },
        {
            title: "Bandwidth",
            dataIndex: "megabytesPerSecond",
            align: "right",
            render: formatMegabytes,
        },
        {
            title: "Average time",
            dataIndex: "averageSeconds",
            align: "right",
            render: formatSeconds,
        },
        {
            title: "Std deviation",
            dataIndex: "stdDeviationSeconds",
            align: "right",
            render: formatSeconds,
        },
        { title: "Samples", dataIndex: "samples", align: "right" },
    ];
}
