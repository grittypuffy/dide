# Development

## Pre-requisites

Ensure you have the following dependencies available on your system:

- [Go (1.21+)](https://go.dev/): Required for Wails 
- [Wails](https://wails.io/docs/gettingstarted/installation): Required for building the application
- npm and Node.js: Required for packaging frontend in the application along with dependency management

## Building

1. Clone the repository:

```shell
git clone https://github.com/grittypuffy/dide
cd dide
```

2. Install dependencies:

```shell
go mod tidy
```

3. Building the application:
    DIDE uses the Wails framework for development and building cross-platform web-based application.
    Follow these instructions for your platform in the root directory of the project:

    - Windows
        ```shell
        wails build -platform windows/amd64
        ```
    - macOS:
        ```shell
        wails build -platform darwin/universal
        ```
    - Linux:
        ```shell
        wails build -platform linux/amd64
        ```

The executable for the corresponding platform must be available in the `build/bin` directory (on Linux and macOS)

4. Run the production build using the following command by changing into `build/bin` directory:

```shell
./dide
```