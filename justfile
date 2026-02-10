# Development with live reload
dev: generate css
    air

# Watch templ files
dev-templ:
    templ generate --watch

# Watch CSS
dev-css:
    tailwindcss -i ./dev/static/styles.css -o ./dev/static/output.css --watch

# Generate templ files
generate:
    templ generate

# Build CSS once
css:
    tailwindcss -i ./dev/static/styles.css -o ./dev/static/output.css --minify

# Combined build
build: generate css
    go build -o bin/garden-dev ./dev

# Run dev server
run:
    go run ./dev/main.go

# Production artifact via Nix
nix-build:
    nix build

nix-run:
    ./result/bin/garden-dev

# Clean
clean:
    rm -rf bin/
    rm -f dev/static/output.css
    find . -name "*_templ.go" -delete
