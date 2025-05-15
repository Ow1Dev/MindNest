# justfile

# Variables
bin := "bin/mindnest"
main := "./cmd/gateway/main.go"

# Default recipe (optional)
default:
  @just --summary

# Build the Go project
build:
  mkdir -p bin
  go build -o {{bin}} {{main}}

# Run the built project (builds it first)
run: build
  ./{{bin}}

# Clean the build artifacts
clean:
  rm -rf bin

# Show available tasks (optional)
help:
  just --summary
