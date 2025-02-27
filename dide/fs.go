package dide

import (
	"encoding/json"
	"os"
	"path/filepath"
	"log"
)

type Node struct {
	Name     string  `json:"name"`
	Folder   bool  `json:"folder"`
	Children []Node `json:"children,omitempty"`
	FilePath string `json:"path"`
}

func traverseDir(path string) ([]Node, error) {
	var nodes []Node

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			childPath := filepath.Join(path, file.Name())
			childNodes, err := traverseDir(childPath)
			if err != nil {
				return nil, err
			}
			nodes = append(nodes, Node{
				Name:     file.Name(),
				Folder:   true,
				Children: childNodes,
				FilePath: childPath,
			})
		} else {
			filePath := filepath.Join(path, file.Name())
			nodes = append(nodes, Node{Name: file.Name(), Folder: false, FilePath: filePath})
		}
	}

	return nodes, nil
}

func FolderTree(directoryName string) (string, error) {
	nodes, err := traverseDir(directoryName)
	if err != nil {
		return "", err
	}

	rootNode := Node{
		Name:     filepath.Base(directoryName),
		Folder: true,
		Children: nodes,
	}

	jsonData, err := json.MarshalIndent(rootNode, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func ReadFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
    if err != nil {
        log.Fatal(err)
    }
	return string(content), nil
}