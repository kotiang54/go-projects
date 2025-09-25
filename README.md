# go-projects

This repository contains a collection of Go programming examples and mini-projects, organized by topic and concept. Each subfolder demonstrates a specific feature or package in Go, with standalone executable files and clear documentation.

## Structure

- **intermediate/env_var/**
  Safe interaction with environment variables, including setting, unsetting, and listing environment variable keys, with attention to security and error handling.

- **intermediate/logging/**
  Demonstrates logging techniques using Go's standard `log` package and the third-party `logrus` library, including custom loggers, log levels, file logging, and structured logging.

- **intermediate/json/**
  Shows JSON encoding and decoding with structs, named and anonymous embedding, marshaling/unmarshaling, handling arrays, and decoding unknown JSON structures into maps.

- **intermediate/io_package/**
  Examples of using Go's `io` package for reading and writing with interfaces, buffers, pipes, and files, including robust error handling and resource management.

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

## Requirements

- Go 1.18 or newer
- For logging examples:
  Install logrus with
  ```sh
  go get github.com/sirupsen/logrus
  ```

## License

This repository is for educational purposes.
