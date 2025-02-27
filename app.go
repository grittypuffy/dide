package main

import (
	"context"
	"dide/dide"
	"fmt"
	"log"
	"strings"

	"github.com/rifaideen/talkative"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	Ollama *talkative.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		Ollama: dide.GetTalkativeClient(),
	}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func (a *App) GetFolderTree(path string) string {
	jsonOutput, err := dide.FolderTree(path)
	if err != nil {
		return ""
	}
	return jsonOutput
}

func (a *App) GetWorkingPath() string {
	selection, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Choose a file",
	})

	if err != nil {
		fmt.Println(err)
	}

	return selection
}

func (a *App) GetFileContent(filePath string) (string, error) {
	content, err := dide.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return content, nil
}

func (a *App) GenerateDocumentation(content string) string {
	llm, err := ollama.New(ollama.WithModel("hf.co/grittypuffy/dide_code_quality:Q4_K_M"))
	if err != nil {
		log.Fatal(err)
	}
	var builder strings.Builder
	prompt := fmt.Sprintf(
		`Give only the necessary documentation along with its corresponding source code without extra explanations in plain text format without HTML or markdown.
Consider the following example input and example output:
Example Input:
def greet(name):
    return f"Hello {name}"
def main():
    return "Hello world"

Example output:
def greet(name):
	"""Greets a user by their name"""
	return f"Hello {name}"

def main():
	"""Returns hello world"""
	return "Hello world"
Do refactoring and inline documentation if needed

Format the given input code in the manner described above: %s`, content)

	ctx := context.Background()
	completion, err := llm.Call(ctx, prompt,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			builder.WriteString(string(chunk))
			return nil
		}),
	)
	result := builder.String()
	if err != nil {
		log.Fatal(err)
	}
	_ = completion
	return result
}

func (a *App) GenerateTests(content string) string {
	llm, err := ollama.New(ollama.WithModel("hf.co/grittypuffy/dide_code_quality:Q4_K_M"))
	if err != nil {
		log.Fatal(err)
	}
	var builder strings.Builder
	prompt := fmt.Sprintf(
		`Give only the test cases along with its corresponding input source code without extra explanations in plain text format without HTML or markdown.
Consider the following example input and example output:
Example Input:
def greet(name):
    return f"Hello {name}"

Example output:
def greet(name):
    return f"Hello {name}"

def test_greet():
    # Test with a simple name
    result = greet("Alice")
    assert result == "Hello Alice", f"Expected 'Hello Alice' but got {result}"

    # Test with another name
    result = greet("Bob")
    assert result == "Hello Bob", f"Expected 'Hello Bob' but got {result}"

    # Test with an empty string as name
    result = greet("")
    assert result == "Hello ", f"Expected 'Hello ' but got {result}"

    # Test with a name containing spaces
    result = greet("Charlie Brown")
    assert result == "Hello Charlie Brown", f"Expected 'Hello Charlie Brown' but got {result}"

    # Test with a long name
    result = greet("A" * 100)
    assert result == "Hello " + "A" * 100, f"Expected 'Hello {'A' * 100}' but got {result}"

    print("All test cases passed!")
Do refactoring and inline documentation if needed

Format the given input code in the manner described above: %s`, content)

	ctx := context.Background()
	completion, err := llm.Call(ctx, prompt,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			builder.WriteString(string(chunk))
			return nil
		}),
	)
	result := builder.String()
	if err != nil {
		log.Fatal(err)
	}
	_ = completion
	return result
}
