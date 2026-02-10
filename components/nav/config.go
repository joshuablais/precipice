// precipice/components/nav/config.go
package nav

import (
	"github.com/a-h/templ"
	"github.com/joshuablais/precipice/components/button"
	"os"
)

type NavItem struct {
	Label string
	Href  string
}

type AuthButton struct {
	Label     string
	Href      string
	Variant   string
	IsPrimary bool
}

type LogoConfig struct {
	DesktopURL string // Logo for main nav
	MobileURL  string // Logo for mobile menu (can be same or different)
	AltText    string
}

type NavConfig struct {
	Items       []NavItem
	AuthButtons []AuthButton
	Logo        LogoConfig
}

// ToButtonArgs converts AuthButton config to button.ButtonArgs
// with additional attributes for context (desktop vs mobile)
func (ab AuthButton) ToButtonArgs(additionalClass string, additionalAttrs templ.Attributes) button.ButtonArgs {
	attrs := templ.Attributes{
		"href": ab.Href,
	}

	// Merge additional attributes
	if additionalAttrs != nil {
		for k, v := range additionalAttrs {
			attrs[k] = v
		}
	}

	return button.ButtonArgs{
		Variant:    ab.Variant,
		Size:       "lg",
		AsChild:    true,
		Class:      additionalClass,
		Attributes: attrs,
	}
}

// getEnvOrDefault returns environment variable value or default
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
				Label:     "Log In",
				Href:      "/login",
				Variant:   "navlink",
				IsPrimary: false,
			},
			{
				Label:     "Sign Up",
				Href:      "/register",
				Variant:   "secondary",
				IsPrimary: true,
			},
		},
		Logo: LogoConfig{
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
