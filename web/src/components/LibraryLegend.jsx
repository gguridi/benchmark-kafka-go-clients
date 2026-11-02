import LibrarySwatch from "@/components/LibrarySwatch";
import { libraryLabel } from "@/theme";

export default function LibraryLegend({ libraries, scheme }) {
    return (
        <ul className="flex flex-wrap gap-x-5 gap-y-2 text-sm text-ink-secondary">
            {libraries.map((library) => (
                <li key={library} className="flex items-center gap-2">
                    <LibrarySwatch library={library} scheme={scheme} />
                    {libraryLabel(library)}
                </li>
            ))}
        </ul>
    );
}
