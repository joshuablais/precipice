package nav

import "os"

type NavItem struct {
	Label string
	Href  string
}

type AuthButton struct {
	Label   string
	Href    string
	Variant string
}

type LogoConfig struct {
	Href    string
	Image   string
	AltText string
}

type NavConfig struct {
	Items       []NavItem
	AuthButtons []AuthButton
	Logo        LogoConfig
}

func GetEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
