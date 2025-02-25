package dide

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Node struct {
	Name     string  `json:"name"`
	Children []Node `json:"children,omitempty"`
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
				Children: childNodes,
			})
		} else {
			nodes = append(nodes, Node{Name: file.Name()})
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
		Children: nodes,
	}

	jsonData, err := json.MarshalIndent(rootNode, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}