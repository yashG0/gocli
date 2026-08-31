# gocli

A personal command-line toolkit written in Go.

`gocli` is a learning project focused on building useful, cross-platform command-line utilities while learning Go through active development.

## Current Features

### System Information

Display basic information about the current system:

```bash
go run ./cmd/gocli info
```

Example output:

```text
System Information
------------------
OS: linux
Architecture: amd64
Hostname: fedora
```

## Available Commands

| Command | Description |
|---------|-------------|
| `info` | Display system information |
| `files` | File and directory utilities |
| `search` | Search for files |

> `files` and `search` are currently under development.

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

Run the CLI:

```bash
go run ./cmd/gocli
```

Run a command:

```bash
go run ./cmd/gocli info
```

## Building

Build the executable:

```bash
go build -o gocli ./cmd/gocli
```

Run the executable:

```bash
./gocli info
```

## Roadmap

- [x] Basic CLI argument handling
- [x] Command routing
- [x] System information
- [ ] File listing
- [ ] File searching
- [ ] Disk usage
- [ ] Process information
- [ ] Network information
- [ ] File backup utility
- [ ] Project statistics
- [ ] Configuration
- [ ] Unit tests
- [ ] Cross-platform support
- [ ] Release binaries

## Goals

- Learn Go through active development
- Build useful command-line tools
- Explore Go's standard library
- Work with files and operating-system APIs
- Learn error handling and testing
- Explore concurrency
- Write clean and maintainable Go code
- Build a real-world portfolio project

## License

MIT
