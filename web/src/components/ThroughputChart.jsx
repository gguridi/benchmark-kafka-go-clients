import {
    Bar,
    BarChart,
    CartesianGrid,
    LabelList,
    ResponsiveContainer,
    Tooltip,
    XAxis,
    YAxis,
} from "recharts";
import ThroughputTooltip from "@/components/ThroughputTooltip";
import { formatCompact, formatMessageCount } from "@/format";
import { CHROME, libraryColor, libraryLabel } from "@/theme";

const CHART_HEIGHT = 280;
const BAR_THICKNESS = 24;
const BAR_GAP = 2;
const DATA_END_RADIUS = [4, 4, 0, 0];

export default function ThroughputChart({ group, libraries, scheme }) {
    const chrome = CHROME[scheme];
    const fastestByCount = new Map(
        group.fastest.map((benchmark) => [benchmark.numMessages, benchmark.library]),
    );
    return (
        <figure aria-label={`Throughput for ${group.messageSize} byte messages`}>
            <ResponsiveContainer width="100%" height={CHART_HEIGHT}>
                <BarChart
                    data={group.rows}
                    barGap={BAR_GAP}
                    barCategoryGap="25%"
                    margin={{ top: 20 }}
                >
                    <CartesianGrid vertical={false} stroke={chrome.grid} />
                    <XAxis
                        dataKey="numMessages"
                        tickFormatter={formatMessageCount}
                        tickLine={false}
                        stroke={chrome.axis}
                        tick={{ fill: chrome.muted, fontSize: 12 }}
                    />
                    <YAxis
                        tickFormatter={formatCompact}
                        axisLine={false}
                        tickLine={false}
                        width={48}
                        tick={{ fill: chrome.muted, fontSize: 12 }}
                    />
                    <Tooltip
                        cursor={{ fill: chrome.grid, fillOpacity: 0.4 }}
                        content={<ThroughputTooltip scheme={scheme} />}
                    />
                    {libraries.map((library) => (
                        <Bar
                            key={library}
                            name={libraryLabel(library)}
                            dataKey={(row) => row[library]?.throughput}
                            fill={libraryColor({ library, scheme })}
                            maxBarSize={BAR_THICKNESS}
                            radius={DATA_END_RADIUS}
                        >
                            <LabelList
                                content={(props) => (
                                    <FastestLabel
                                        {...props}
                                        isFastest={
                                            fastestByCount.get(
                                                group.rows[props.index]?.numMessages,
                                            ) === library
                                        }
                                        color={chrome.muted}
                                    />
                                )}
                            />
                        </Bar>
                    ))}
                </BarChart>
            </ResponsiveContainer>
        </figure>
    );
}

function FastestLabel({ x, y, width, value, isFastest, color }) {
    if (!isFastest || value === undefined) {
        return null;
    }
    return (
        <text x={x + width / 2} y={y - 6} textAnchor="middle" fontSize={12} fill={color}>
            {formatCompact(value)}
        </text>
    );
}
