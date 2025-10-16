# GoLevelUI: A Web Admin Tool for LevelDB

A simple, embeddable web GUI for managing LevelDB databases, written in Go with a Vue.js frontend.

## Features

-   **CRUD Operations**: Create, read, update, and delete key-value pairs.
-   **Embeddable**: Integrate into an existing Go application or run as a standalone server.
-   **Multiple DB Support**: Manage multiple LevelDB instances in one interface.

## Usage

### 1. As a Go Library

Embed GoLevelUI directly into your application. It will start a web server on port `8080` by default.

```go
package main

import (
	"github.com/liuzl/golevelui"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

func main() {
	// Open your LevelDB instance
	db, err := leveldb.OpenFile("/tmp/example_db", nil)
	if err != nil {
		log.Fatalf("Failed to open leveldb: %v", err)
	}

	// Register the DB instance with a unique name and start the admin server
	err = golevelui.GetLevelAdmin().Register(db, "My App DB").Start()
	if err != nil {
		log.Fatalf("Failed to start GoLevelUI: %v", err)
	}
	
	// Keep your application running
	select{}
}
```

### 2. As a Standalone Executable

Build and run the tool from the command line, passing the paths to your LevelDB directories.

```bash
# Build the executable
go build -o golevelui_server ./cmd/main.go

# Run the server, specifying one or more database paths
./golevelui_server --path /path/to/db1 --path /path/to/db2
```

## Accessing the UI

Once the server is running, open your browser and navigate to:
**http://12.0.0.1:8080/golevelui/static/**