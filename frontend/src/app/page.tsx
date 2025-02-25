'use client'
import { Editor } from "@monaco-editor/react"
import { useState } from "react";
import { GetWorkingPath } from "../../wailsjs/go/main/App";
import { GetFolderTree } from "../../wailsjs/go/main/App";
import Entry from "@/components/FolderTree";

export default function Home() {
    const [workingPath, setWorkingPath] = useState< string |null >(null);
    const [folderTree, setFolderTree] = useState(null);

    const selectFolder = async () => {
        const filePath = await GetWorkingPath();
        setWorkingPath(filePath);
        const folderTreeRes = await GetFolderTree(filePath);
        console.log(folderTreeRes);
        if (folderTreeRes !== "") {
            setFolderTree(JSON.parse(folderTreeRes))
        }
    }

    return (
        <div className="grid grid-cols-8 items-center space-y-6 max-w-full max-h-full">
            <div className="h-full w-full flex flex-col items-center justify-items-center py-4 text-slate-800">
                {workingPath && <p>{workingPath.split("/").pop()}</p>}
                <button className="px-2 py-1 bg-sky-400 border-2 border-sky-500 border-solid text-sm font-semibold" onClick={selectFolder}>Select Folder</button>
                {folderTree && <Entry entry={folderTree} depth={1}></Entry>}
            </div>
            <Editor
                className="col-start-4"
                height="100vh"
                width="100vw"
                defaultLanguage="python"
                defaultValue={`def main():
    print("Hello World")
main()
`}
            />
        </div>
    )
}
