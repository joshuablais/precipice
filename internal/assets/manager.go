// internal/assets/manager.go
// A helper function for defining asset locations

package assets

import (
	"log"
	"os"
	"strings"
)

type Manager struct {
	baseURL    string
	localPath  string
	production bool
}

func NewManager() *Manager {
	production := os.Getenv("ENVIRONMENT") == "production"
	cdnURL := os.Getenv("CDN_URL")

	// Ensure protocol is present in production
	if production && cdnURL != "" && !strings.HasPrefix(cdnURL, "http") {
		cdnURL = "https://" + cdnURL
	}

	m := &Manager{
		baseURL:    cdnURL,
		localPath:  "/assets",
		production: production,
	}

	// Log on initialization so you know what mode you're in
	if production {
		log.Printf("Assets: Production mode - CDN: %s", cdnURL)
	} else {
		log.Printf("Assets: Development mode - Local: %s", m.localPath)
	}

	return m
}

func (m *Manager) URL(assetPath string) string {
	// Ensure assetPath starts with /
	if !strings.HasPrefix(assetPath, "/") {
		assetPath = "/" + assetPath
	}

	if m.production && m.baseURL != "" {
		return m.baseURL + assetPath
	}

	// Development: concatenate, do NOT use filepath.Join
	return m.localPath + assetPath
}
