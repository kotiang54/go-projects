# go-projects

This repository contains a collection of Go programming examples and mini-projects, organized by topic and concept. Each subfolder demonstrates a specific feature or package in Go, with standalone executable files and clear documentation.

## Structure

- **basics/**
  Foundational Go concepts: variables, functions, control flow, slices, maps, and basic error handling.

- **intermediate/**
  Examples covering:
  - Safe interaction with environment variables (setting, unsetting, listing keys, security, error handling)
  - Logging techniques with Go's standard `log` package and third-party `logrus` (custom loggers, log levels, file logging, structured logging)
  - JSON encoding/decoding with structs, embedding, marshaling/unmarshaling, arrays, and decoding unknown structures
  - Using Go's `io` package for reading/writing with interfaces, buffers, pipes, files, robust error handling, and resource management

- **advanced/**
  Advanced Go topics such as concurrency (goroutines, channels), context usage, and custom error types.

- **games/**
  Simple game implementations in Go, demonstrating application structure, user input, and basic game logic.

## Usage

Each folder contains a standalone Go file with a `main()` function. To run an example:

```sh
cd intermediate/env_var
go run env_var.go
```

Replace the folder and filename as needed for other examples.

## Highlights

- Secure handling of sensitive data and environment variables.
- Structured and robust error handling.
- Use of Go interfaces (`io.Reader`, `io.Writer`, `io.Closer`) for flexible I/O.
- Demonstrations of third-party libraries for advanced logging.
- Practical examples for marshaling and unmarshaling JSON data.
- Concurrency and context management examples.
- Game logic and user interaction samples.

## Requirements

- Go 1.18 or newer
- For logging examples:
  Install logrus with
  ```sh
  go get github.com/sirupsen/logrus
  ```

## License

This repository is for educational purposes.
