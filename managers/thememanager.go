package managers

import (
	"fmt"
	"slices"
	"strings"
	d "tuidoo/dictionaries"

	e "tuidoo/entities"
)

type ThemeManager struct {
	Themes       map[string]e.Theme
	CurrentTheme e.Theme
}

func NewThemeManager() *ThemeManager {
	tm := &ThemeManager{
		Themes:       d.Themes,
		CurrentTheme: d.Themes["dark"],
	}
	return tm
}

func (tm *ThemeManager) SetTheme(id string) error {
	if _, exists := tm.Themes[id]; !exists {
		return fmt.Errorf("theme '%s' not found", id)
	}
	tm.CurrentTheme = tm.Themes[id]
	return nil
}

func (tm *ThemeManager) GetCurrentTheme() e.Theme {
	return tm.CurrentTheme
}

func (tm *ThemeManager) GetThemeNames() []string {
	names := make([]string, 0, len(tm.Themes))
	for id := range tm.Themes {
		names = append(names, id)
	}
	return names
}

func (tm *ThemeManager) GetCurrentThemeIndex() int {
	themes := tm.GetThemeNames()
	index := slices.Index(themes, tm.CurrentTheme.ID)
	if index == -1 {
		return 0
	}
	return index
}

func (tm *ThemeManager) CreateHeader() string {
	lines := []string{
		"████████╗ ██╗   ██╗ ██╗ ██████╗   ██████╗   ██████╗ ",
		"╚══██╔══╝ ██║   ██║ ██║ ██╔══██╗ ██╔═══██╗ ██╔═══██╗",
		"   ██║    ██║   ██║ ██║ ██║  ██║ ██║   ██║ ██║   ██║",
		"   ██║    ██║   ██║ ██║ ██║  ██║ ██║   ██║ ██║   ██║",
		"   ██║    ╚██████╔╝ ██║ ██████╔╝ ╚██████╔╝ ╚██████╔╝",
		"   ╚═╝     ╚═════╝  ╚═╝ ╚═════╝   ╚═════╝   ╚═════╝ ",
	}

	// Define the gradient colors matching your logo (left to right: green → cyan → blue)
	// Gradient: green → cyan → blue (matching logo)
	gradientColors := []struct{ r, g, b uint8 }{
		{0, 255, 0},   // #00FF00 bright green
		{0, 255, 102}, // #00FF66
		{0, 255, 204}, // #00FFCC
		{0, 204, 255}, // #00CCFF
		{0, 153, 255}, // #0099FF
		{0, 102, 255}, // #0066FF blue
	}

	var result strings.Builder

	for _, line := range lines {
		lineLen := len(line)

		for charIdx, char := range line {
			progress := float64(charIdx) / float64(lineLen)
			colorIdx := progress * float64(len(gradientColors)-1)

			idx1 := int(colorIdx)
			idx2 := idx1 + 1
			if idx2 >= len(gradientColors) {
				idx2 = len(gradientColors) - 1
			}

			t := colorIdx - float64(idx1)
			color1 := gradientColors[idx1]
			color2 := gradientColors[idx2]

			r := uint8(float64(color1.r)*(1-t) + float64(color2.r)*t)
			g := uint8(float64(color1.g)*(1-t) + float64(color2.g)*t)
			b := uint8(float64(color1.b)*(1-t) + float64(color2.b)*t)

			result.WriteString(fmt.Sprintf("[#%02x%02x%02x]%c[-]", r, g, b, char))
		}
		result.WriteString("\n")
	}

	return result.String()
}
