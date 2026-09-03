# gocli

A personal command-line toolkit written in Go.

`gocli` is a learning project focused on building useful, cross-platform
command-line utilities while learning Go through hands-on development.

The project primarily uses Go's standard library.

## Current Features

### System Information

Display basic information about the current system:

```bash
./gocli info
```

Example:

```text
System Information
------------------
OS: linux
Architecture: amd64
Hostname: fedora
CPU Cores: 4
```

### File Listing

List files and directories from a specified path:

```bash
./gocli files .
```

### Recursive File Search

Search for a file recursively through a directory and its subdirectories:

```bash
./gocli search main.go
```

By default, the current directory is searched.

You can also specify a search location:

```bash
./gocli search main.go /tmp
```

Example:

```text
Found: cmd/gocli/main.go
Total: 1 found!
```

## Available Commands

| Command | Description |
|---------|-------------|
| `info` | Display system information |
| `files <path>` | List files and directories |
| `search <filename> [path]` | Recursively search for a file |

## Project Structure

```text
gocli/
├── cmd/
│   └── gocli/
│       └── main.go
├── internal/
├── go.mod
├── README.md
└── .gitignore
```

## Running

Make sure Go is installed.

Run the project directly:

```bash
go run ./cmd/gocli
```

Run a command:

```bash
go run ./cmd/gocli info
go run ./cmd/gocli files .
go run ./cmd/gocli search main.go
```

## Building

Build the executable:

```bash
go build -o gocli ./cmd/gocli
```

Run the executable:

```bash
./gocli info
./gocli files .
./gocli search main.go
```

## Roadmap

- [x] Basic CLI argument handling
- [x] Command routing
- [x] System information
- [x] File listing
- [x] Recursive file searching
- [x] Search result counting
- [x] Git and GitHub integration
- [ ] Improve file listing output
- [ ] Improve search output
- [ ] Disk usage
- [ ] Process information
- [ ] Network information
- [ ] File backup utility
- [ ] Project statistics
- [ ] Configuration
- [ ] Unit tests
- [ ] Cross-platform testing
- [ ] Release binaries

## Goals

- Learn Go through active development
- Build useful command-line tools
- Explore Go's standard library
- Work with files and operating-system APIs
- Learn error handling
- Learn recursion and filesystem traversal
- Explore concurrency
- Write clean and maintainable Go code
- Build a real-world portfolio project

## Development Philosophy

This project is built incrementally rather than following a tutorial.

Features are implemented, tested, committed to Git, and pushed to GitHub
as the project evolves.

The goal is to understand how things work rather than simply copy
existing implementations.

## License

MIT
