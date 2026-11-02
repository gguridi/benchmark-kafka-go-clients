import { ConfigProvider, Segmented, theme } from "antd";
import { useSyncExternalStore } from "react";
import { CONSUMER, PRODUCER } from "@/benchmarks";
import BenchmarkPage from "@/components/BenchmarkPage";
import { CHROME, DARK, useColorScheme } from "@/theme";

const REPOSITORY_URL = "https://github.com/gguridi/benchmark-kafka-go-clients";

const PAGES = [
    { label: "Producers", value: PRODUCER, hash: "#/producers" },
    { label: "Consumers", value: CONSUMER, hash: "#/consumers" },
];

export default function App() {
    const scheme = useColorScheme();
    const page = useCurrentPage();
    return (
        <ConfigProvider
            theme={{
                algorithm: scheme === DARK ? theme.darkAlgorithm : theme.defaultAlgorithm,
                token: {
                    fontFamily: "system-ui, -apple-system, 'Segoe UI', sans-serif",
                    colorBgContainer: CHROME[scheme].surface,
                },
            }}
        >
            <main className="mx-auto flex max-w-6xl flex-col gap-6 px-4 py-8 sm:px-6">
                <header className="flex flex-col gap-4 sm:flex-row sm:items-end sm:justify-between">
                    <div className="flex flex-col gap-1">
                        <h1 className="text-2xl font-semibold">Kafka Go clients benchmark</h1>
                        <p className="max-w-2xl text-sm text-ink-secondary">
                            Throughput of each client, in messages per second: higher is better.
                            Every configuration runs all clients one after another on the same
                            GitHub Actions runner and Kafka broker.{" "}
                            <a className="underline" href={REPOSITORY_URL}>
                                Source and methodology
                            </a>
                        </p>
                    </div>
                    <Segmented
                        options={PAGES}
                        value={page.value}
                        onChange={(value) => navigateTo(value)}
                    />
                </header>
                <BenchmarkPage type={page.value} scheme={scheme} />
            </main>
        </ConfigProvider>
    );
}

function useCurrentPage() {
    const hash = useSyncExternalStore(subscribeToHash, () => window.location.hash);
    return PAGES.find((page) => page.hash === hash) ?? PAGES[0];
}

function subscribeToHash(onChange) {
    window.addEventListener("hashchange", onChange);
    return () => window.removeEventListener("hashchange", onChange);
}

function navigateTo(value) {
    window.location.hash = PAGES.find((page) => page.value === value).hash;
}
