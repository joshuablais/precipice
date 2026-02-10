# Generate templ files
generate:
    templ generate

# Run once
run: generate
    go run main.go

# Live reload (actual hot reloading)
dev:
    air

# Build for production
build: generate
    go build -o bin/precipice .

# Clean generated files
clean:
    rm -rf bin/ tmp/ **/*_templ.go
