.PHONY: build build-all run seed reset clean install help

# Build main TUI application
build:
	go build -o bin/tuidoo cmd/tuidoo/main.go

# Build all tools
build-all:
	go build -o bin/tuidoo cmd/tuidoo/main.go
	go build -o bin/tuidoo-seed cmd/seed/main.go
	go build -o bin/tuidoo-reset cmd/reset/main.go
	go build -o bin/tuidoo-clean cmd/clean/main.go

# Run TUI application
run:
	go run tui/app/main.go

# Seed database
seed:
	go run cli/seed/main.go

# Reset database
reset:
	go run cli/reset/main.go

# Clean database
clean:
	go run cli/clean/main.go

# Install to system
install: build-all
	cp bin/tuidoo $(GOPATH)/bin/
	cp bin/tuidoo-seed $(GOPATH)/bin/
	cp bin/tuidoo-reset $(GOPATH)/bin/
	cp bin/tuidoo-clean $(GOPATH)/bin/

help:
	@echo "TUIDOO - Build Commands"
	@echo ""
	@echo "  make build       Build main TUI application"
	@echo "  make build-all   Build all binaries"
	@echo "  make run         Run TUI application"
	@echo "  make seed        Seed database"
	@echo "  make reset       Reset database"
	@echo "  make clean       Clean database"
	@echo "  make install     Install to system"
