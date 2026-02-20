// components/titles/types.go

package titles

import (
	"fmt"
	"strings"
)

type Alignment string
type Background string
type Size string

const (
	AlignCenter Alignment = "center"
	AlignLeft   Alignment = "left"

	BgDark     Background = "dark"
	BgGradient Background = "gradient"
	BgImage    Background = "image"
	BgLight    Background = "light"

	SizeSm Size = "sm"
	SizeMd Size = "md"
	SizeLg Size = "lg"
)

type Announcement struct {
	Text     string
	LinkText string
	Href     string
}

type TitleSectionProps struct {
	Title           string
	Subtitle        string
	Alignment       Alignment
	Background      Background
	BackgroundImage string
	Size            Size
	Announcement    *Announcement
}

func (p TitleSectionProps) classes() string {
	align := p.Alignment
	if align == "" {
		align = AlignCenter
	}

	size := p.Size
	if size == "" {
		size = SizeMd
	}

	bg := p.Background
	if bg == "" {
		bg = BgDark
	}

	return strings.Join([]string{
		"title-section",
		fmt.Sprintf("title-section--%s", align),
		fmt.Sprintf("title-section--%s", size),
		fmt.Sprintf("title-section--%s", bg),
	}, " ")
}

func (p TitleSectionProps) style() string {
	if p.Background == BgImage && p.BackgroundImage != "" {
		return fmt.Sprintf("--ts-bg-image: url(%s)", p.BackgroundImage)
	}
	return ""
}
