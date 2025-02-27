'use client'

import { BreadcrumbItem, Breadcrumbs, Button } from "@heroui/react";
import { Editor } from "@monaco-editor/react";
import { useState } from "react";
import { GetWorkingPath, GetFolderTree, GenerateDocumentation, GenerateTests } from "../../wailsjs/go/main/App";
import Entry from "./FolderTree";
import { Folder } from "lucide-react";

export default function App() {
    const [workingPath, setWorkingPath] = useState<string | null>(null);
    const [folderTree, setFolderTree] = useState(null);
    const [editorContent, setEditorContent] = useState("");

    const updateEditorContent = (content: string) => {
        setEditorContent(content);
    }
    const generateDocumentation = async (sourceCode: string) => {
        const documentedCode = await GenerateDocumentation(sourceCode);
        setEditorContent(documentedCode);
    }

    const generateTests = async (sourceCode: string) => {
        const testCode = await GenerateTests(sourceCode);
        setEditorContent(testCode);
    }

    const selectFolder = async () => {
        const filePath = await GetWorkingPath();
        setWorkingPath(filePath);
        const folderTreeRes = await GetFolderTree(filePath);
        if (folderTreeRes !== "") {
            setFolderTree(JSON.parse(folderTreeRes))
        }
    }

    return (
        <div className="grid grid-cols-8 space-y-6 h-screen w-screen">
            <div className="flex flex-col space-y-2 space-x-2 items-start py-4 text-slate-800 bg-slate-100">
                <Button color="success" endContent={<Folder />} onPress={selectFolder}>Select Folder</Button>
                {workingPath && (<Breadcrumbs variant="solid">
                    <BreadcrumbItem>{workingPath.split("/").pop()}</BreadcrumbItem>
                </Breadcrumbs>)}
                {folderTree && <Entry entry={folderTree} depth={1} updateEditorContent={updateEditorContent}></Entry>}
            </div>
            <div className="grid grid-rows-10 items-start col-start-2 w-screen h-screen">
                {
                    editorContent && (<div className="px-6 flex gap-4 items-start">
                        <Button color="secondary" radius="md" onPress={async () => await generateDocumentation(editorContent)}>Generate documentation</Button>
                        <Button color="secondary" radius="md" onPress={async () => await generateTests(editorContent)}>Generate Test Cases</Button>
                        <Button color="secondary" radius="md">Generate Code Completion</Button>
                    </div>)
                }
                <Editor
                    className="row-start-2"
                    width="87.5vw"
                    height="90vh"
                    defaultLanguage="python"
                    value={editorContent}
                />
            </div>
        </div>)
}