const COMPACT = new Intl.NumberFormat("en", { notation: "compact", maximumFractionDigits: 1 });
const DECIMAL = new Intl.NumberFormat("en", { maximumFractionDigits: 1 });
const SECONDS = new Intl.NumberFormat("en", { maximumFractionDigits: 3 });

export const formatCompact = (value) => COMPACT.format(value);

export const formatThroughput = (value) => `${COMPACT.format(value)} msg/s`;

export const formatMegabytes = (value) => `${DECIMAL.format(value)} MB/s`;

export const formatSeconds = (value) => `${SECONDS.format(value)} s`;

export const formatMessageSize = (bytes) => `${DECIMAL.format(bytes / 1000)} KB`;

export const formatMessageCount = (count) => `${COMPACT.format(count)} messages`;
