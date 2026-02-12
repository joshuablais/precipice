# Generate templ files
generate:
    templ generate

# Run once
run: generate
    go run cmd/main.go

# Live reload (actual hot reloading)
dev:
    air & sleep 1 && chromium --new-window http://localhost:3000

# Build for production
build: generate
    go build -o bin/precipice .

# Clean generated files
clean:
    rm -rf bin/ tmp/ **/*_templ.go

# Build CSS
cssbuild:
    esbuild static/main.css --bundle --minify --outfile=dist/css/main.css
