// precipice/components/nav/config.go
// precipice/components/nav/config.go
package nav

import (
	"os"
)

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
	Href       string // Where the logo links to
	DesktopURL string // Logo image for main nav
	MobileURL  string // Logo image for mobile menu
	AltText    string
}

type NavConfig struct {
	Items       []NavItem
	AuthButtons []AuthButton
	Logo        LogoConfig
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func GetNavConfig() NavConfig {
	return NavConfig{
		Items: []NavItem{
			{Label: "For Agents", Href: "/agents"},
			{Label: "For Brokers", Href: "/brokers"},
			{Label: "Features", Href: "#features"},
			{Label: "Pricing", Href: "#pricing"},
			{Label: "Contact", Href: "/contact"},
		},
		AuthButtons: []AuthButton{
			{
				Label:   "Log In",
				Href:    "/login",
				Variant: "ghost",
			},
			{
				Label:   "Sign Up",
				Href:    "/register",
				Variant: "primary",
			},
		},
		Logo: LogoConfig{
			Href: "/",
			DesktopURL: getEnvOrDefault(
				"NAV_LOGO_DESKTOP_URL",
				"https://cella.b-cdn.net/Revere/revere_logo2.png",
			),
			MobileURL: getEnvOrDefault(
				"NAV_LOGO_MOBILE_URL",
				"https://cella.b-cdn.net/Revere/revere_logo2.png",
			),
			AltText: getEnvOrDefault("APP_NAME", "Logo"),
		},
	}
}
