import { libraryColor } from "@/theme";

export default function LibrarySwatch({ library, scheme }) {
    return (
        <span
            aria-hidden="true"
            className="inline-block size-2.5 shrink-0 rounded-sm"
            style={{ background: libraryColor({ library, scheme }) }}
        />
    );
}
