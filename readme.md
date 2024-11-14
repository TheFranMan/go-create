# Go Create
A command-line tool for generating new Go projects based on a predefined template. This tool saves time by setting up a blank project structure according to Go best practices.

## Prerequisites
Ensure you have [Go installed](https://go.dev/doc/install) on your system.

## Installation
To build the go-create binary, run:

```bash
go build -o go-create
```
This will compile the binary, which you can then use to create new Go projects.

## Usage
To create a new project, use the following command:

```bash
./go-create <name> <version>
```

### Parameters
   * `name`: The name of the new application.
   * `version`: The Go version to use, e.g., "1.23.0".

### Example
```bash
./go-create myapp 1.23.0
```

This command generates a new project named myapp configured to use Go version 1.23.0.

## Project Structure
A new project created by go-create will have a structure similar to:

```go
myapp/
├── application
|   └── application.go
├── common
|   ├── config.go
|   └── config_test.go
├── server
|   └── server.go
├── go.mod
├── main.go
├── .env
└── Dockerfile
```

## Testing
To run the tests, execute:

```bash
go test ./...
```

This runs all tests within the project recursively.

