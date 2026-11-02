import { useSyncExternalStore } from "react";

export const LIGHT = "light";
export const DARK = "dark";

const DARK_QUERY = "(prefers-color-scheme: dark)";

const CLIENTS = {
    confluent: { label: "confluent-kafka-go", color: { [LIGHT]: "#2a78d6", [DARK]: "#3987e5" } },
    sarama: { label: "Sarama", color: { [LIGHT]: "#eb6834", [DARK]: "#d95926" } },
    kafkago: { label: "kafka-go", color: { [LIGHT]: "#1baf7a", [DARK]: "#199e70" } },
    franzgo: { label: "franz-go", color: { [LIGHT]: "#eda100", [DARK]: "#c98500" } },
};

const LIBRARY_CLIENT = {
    confluent: CLIENTS.confluent,
    "confluent-poll": CLIENTS.confluent,
    sarama: CLIENTS.sarama,
    kafkago: CLIENTS.kafkago,
    franzgo: CLIENTS.franzgo,
};

export const CHROME = {
    [LIGHT]: { grid: "#e1e0d9", axis: "#c3c2b7", muted: "#898781", surface: "#fcfcfb" },
    [DARK]: { grid: "#2c2c2a", axis: "#383835", muted: "#898781", surface: "#1a1a19" },
};

export const libraryLabel = (library) => LIBRARY_CLIENT[library].label;

export const libraryColor = ({ library, scheme }) => LIBRARY_CLIENT[library].color[scheme];

export function useColorScheme() {
    return useSyncExternalStore(subscribeToColorScheme, currentColorScheme);
}

function subscribeToColorScheme(onChange) {
    const query = window.matchMedia(DARK_QUERY);
    query.addEventListener("change", onChange);
    return () => query.removeEventListener("change", onChange);
}

function currentColorScheme() {
    return window.matchMedia(DARK_QUERY).matches ? DARK : LIGHT;
}
