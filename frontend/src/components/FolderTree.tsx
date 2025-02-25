import { FC, useState } from "react";

export type TFiles = {
    name: string;
    children?: TFiles[]
}

export type EntryProps = {
    entry: TFiles;
    depth: number;
}

export default function Entry({ entry, depth }: EntryProps): ReturnType<FC> {
    const [isExpanded, setIsExpanded] = useState<boolean>(false);
    return (
        <>
            <button onClick={() => setIsExpanded(prev => !prev)} className="flex flex-row">
                {entry.children && (
                    <div className="w-4 pl-1 pr-1 text-xs">
                        {isExpanded ? "-" : "+"}
                    </div>
                )}
                <span style={{ paddingLeft: entry.children ? "" : "8px" }} >{entry.name}</span>
            </button>
            <div className="border-l-2 border-solid border-black m-1 text-xs">
                {isExpanded && <div className="pl-2 text-xs">
                    {entry.children?.map((child) => (
                        <Entry entry={child} depth={depth + 1} key={child.name} />
                    ))}
                </div>}
            </div>
        </>
    )
}