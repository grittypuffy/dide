import { FC, useState } from "react";
import { GetFileContent } from "../../wailsjs/go/main/App";

export type TFiles = {
    name: string;
    children?: TFiles[];
    path: string;
    folder: boolean;
}

export type EntryProps = {
    entry: TFiles;
    depth: number;
    updateEditorContent: (content: string) => void;
}

export default function Entry({ entry, depth, updateEditorContent }: EntryProps): ReturnType<FC> {
    const [isExpanded, setIsExpanded] = useState<boolean>(false);
    const handleFile = async (filePath: string, folder: boolean) => {
        if (!folder) {
            try {
                const content = await GetFileContent(filePath);
                updateEditorContent(content);
            } catch (err) {
                console.log(err);
            }
        }
    }
    return (
        <>
            <button onClick={() => {
                setIsExpanded(prev => !prev);
                handleFile(entry.path, entry.folder)
            }} className="flex flex-row">
                {entry.folder && (
                    <div className="w-4 pl-1 pr-1 text-xs">
                        {isExpanded ? "-" : "+"}
                    </div>
                )}
                <span key={entry.path} style={{ paddingLeft: entry.children ? "" : "8px" }} >{entry.name}</span>
            </button>
            <div className="border-l-2 border-solid border-black m-1 text-xs">
                {isExpanded && <div className="pl-2 text-xs">
                    {entry.children?.map((child) => (
                        <Entry entry={child} depth={depth + 1} key={child.path} updateEditorContent={updateEditorContent}/>
                    ))}
                </div>}
            </div>
        </>
    )
}