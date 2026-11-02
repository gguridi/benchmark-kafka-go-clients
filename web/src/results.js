const RESULT_FILES = import.meta.glob("./results/*.json", { eager: true, import: "default" });

export const RESULTS = Object.values(RESULT_FILES);
