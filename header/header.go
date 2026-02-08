package main

import (
	"fmt"
	"strings"
)

type RGB struct {
	R, G, B int
}

var lines = []string{
	"████████╗██╗   ██╗██╗██████╗  ██████╗  ██████╗ ",
	"╚══██╔══╝██║   ██║██║██╔══██╗██╔═══██╗██╔═══██╗",
	"   ██║   ██║   ██║██║██║  ██║██║   ██║██║   ██║",
	"   ██║   ██║   ██║██║██║  ██║██║   ██║██║   ██║",
	"   ██║   ╚██████╔╝██║██████╔╝╚██████╔╝╚██████╔╝",
	"   ╚═╝    ╚═════╝ ╚═╝╚═════╝  ╚═════╝  ╚═════╝ ",
}

func CreateDiagonalGradient() string {
	var result strings.Builder
	numLines := len(lines)
	maxLen := 51

	for lineIdx, line := range lines {
		runes := []rune(line)

		for charIdx, r := range runes {
			// Diagonal position: top-right = 1.0 (blue), bottom-left = 0.0 (green)
			normalizedChar := float64(charIdx) / float64(maxLen)
			normalizedLine := float64(lineIdx) / float64(numLines-1)

			// Diagonal from top-right to bottom-left
			diagonalPos := (normalizedChar + (1.0 - normalizedLine)) / 2.0

			color := getGreenCyanBlueColor(diagonalPos)
			result.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%c", color.R, color.G, color.B, r))
		}
		result.WriteString("\n")
	}

	result.WriteString("\x1b[0m") // Reset color
	return result.String()
}

func getGreenCyanBlueColor(t float64) RGB {
	// t = 0.0: Bottom-left (Bright Green #00FF00)
	// t = 0.5: Middle diagonal (Cyan #00FFFF)
	// t = 1.0: Top-right (Blue #0099FF)

	var r, g, b int

	if t < 0.5 {
		// Green (#00FF00) to Cyan (#00FFFF)
		progress := t * 2 // 0 to 1
		r = 0
		g = 255
		b = int(progress * 255)
	} else {
		// Cyan (#00FFFF) to Blue (#0099FF)
		progress := (t - 0.5) * 2 // 0 to 1
		r = 0
		g = int(255 - (progress * 102)) // 255 down to 153 (0x99)
		b = 255
	}

	return RGB{R: r, G: g, B: b}
}

var (
	// CyberpunkNeon Cyberpunk Neon (Hot Pink to Cyan)
	CyberpunkNeon = []RGB{
		{R: 255, G: 0, B: 255}, // Magenta
		{R: 255, G: 0, B: 200}, // Hot Pink
		{R: 255, G: 0, B: 150}, // Pink
		{R: 200, G: 0, B: 255}, // Purple
		{R: 100, G: 0, B: 255}, // Blue-Purple
		{R: 0, G: 255, B: 255}, // Cyan
	}

	// Fire Heat (Red to Yellow)
	Fire = []RGB{
		{R: 255, G: 0, B: 0},   // Red
		{R: 255, G: 50, B: 0},  // Red-Orange
		{R: 255, G: 100, B: 0}, // Orange
		{R: 255, G: 150, B: 0}, // Orange-Yellow
		{R: 255, G: 200, B: 0}, // Yellow-Orange
		{R: 255, G: 255, B: 0}, // Yellow
	}

	// Ocean Depths (Deep Blue to Turquoise)
	Ocean = []RGB{
		{R: 0, G: 51, B: 102},    // Deep Ocean Blue
		{R: 0, G: 102, B: 153},   // Ocean Blue
		{R: 0, G: 153, B: 204},   // Light Ocean Blue
		{R: 0, G: 180, B: 220},   // Sky Blue
		{R: 64, G: 224, B: 208},  // Turquoise
		{R: 127, G: 255, B: 212}, // Aquamarine
	}

	// Matrix Green (Dark to Bright Green)
	Matrix = []RGB{
		{R: 0, G: 80, B: 0},      // Dark Green
		{R: 0, G: 120, B: 0},     // Forest Green
		{R: 0, G: 160, B: 0},     // Medium Green
		{R: 0, G: 200, B: 0},     // Bright Green
		{R: 0, G: 240, B: 0},     // Very Bright Green
		{R: 100, G: 255, B: 100}, // Neon Green
	}

	// Sunset (Orange to Purple)
	Sunset = []RGB{
		{R: 255, G: 140, B: 0},  // Dark Orange
		{R: 255, G: 100, B: 50}, // Coral
		{R: 255, G: 69, B: 100}, // Salmon
		{R: 220, G: 20, B: 150}, // Crimson-Pink
		{R: 180, G: 0, B: 200},  // Purple-Pink
		{R: 138, G: 43, B: 226}, // Blue Violet
	}

	// Arctic (Ice Blue to White)
	Arctic = []RGB{
		{R: 100, G: 150, B: 200}, // Ice Blue
		{R: 130, G: 180, B: 220}, // Light Ice
		{R: 160, G: 200, B: 235}, // Pale Blue
		{R: 190, G: 220, B: 245}, // Very Pale Blue
		{R: 220, G: 240, B: 255}, // Almost White Blue
		{R: 240, G: 250, B: 255}, // Ice White
	}

	// RetroAmber Retro Terminal (Amber/Orange Monochrome)
	RetroAmber = []RGB{
		{R: 255, G: 100, B: 0},   // Dark Amber
		{R: 255, G: 130, B: 0},   // Amber
		{R: 255, G: 160, B: 0},   // Light Amber
		{R: 255, G: 180, B: 50},  // Bright Amber
		{R: 255, G: 200, B: 100}, // Very Bright Amber
		{R: 255, G: 220, B: 150}, // Pale Amber
	}

	// Synthwave (Purple, Pink, Blue)
	Synthwave = []RGB{
		{R: 255, G: 0, B: 128},   // Hot Pink
		{R: 255, G: 0, B: 200},   // Pink
		{R: 220, G: 0, B: 255},   // Pink-Purple
		{R: 150, G: 0, B: 255},   // Purple
		{R: 100, G: 100, B: 255}, // Blue-Purple
		{R: 0, G: 200, B: 255},   // Electric Blue
	}

	// Forest (Green to Yellow-Green)
	Forest = []RGB{
		{R: 0, G: 100, B: 0},     // Dark Forest Green
		{R: 34, G: 139, B: 34},   // Forest Green
		{R: 50, G: 180, B: 50},   // Medium Green
		{R: 100, G: 200, B: 50},  // Lime Green
		{R: 150, G: 220, B: 50},  // Yellow-Green
		{R: 180, G: 240, B: 100}, // Light Yellow-Green
	}

	// Dracula Theme Inspired
	Dracula = []RGB{
		{R: 189, G: 147, B: 249}, // Purple
		{R: 139, G: 233, B: 253}, // Cyan
		{R: 80, G: 250, B: 123},  // Green
		{R: 255, G: 184, B: 108}, // Orange
		{R: 255, G: 121, B: 198}, // Pink
		{R: 241, G: 250, B: 140}, // Yellow
	}

	// Nord Theme Inspired
	Nord = []RGB{
		{R: 143, G: 188, B: 187}, // Frost Green
		{R: 136, G: 192, B: 208}, // Frost Blue
		{R: 129, G: 161, B: 193}, // Light Blue
		{R: 94, G: 129, B: 172},  // Medium Blue
		{R: 81, G: 163, B: 163},  // Teal
		{R: 163, G: 190, B: 140}, // Green
	}

	// Vaporwave (Pink to Blue Pastel)
	Vaporwave = []RGB{
		{R: 255, G: 113, B: 206}, // Pink
		{R: 255, G: 140, B: 220}, // Light Pink
		{R: 200, G: 150, B: 255}, // Lavender
		{R: 150, G: 180, B: 255}, // Light Purple
		{R: 120, G: 200, B: 255}, // Sky Blue
		{R: 100, G: 220, B: 255}, // Cyan Blue
	}

	// MonoBlue Monochrome Blue (Dark to Light)
	MonoBlue = []RGB{
		{R: 0, G: 50, B: 150},    // Navy
		{R: 0, G: 100, B: 200},   // Royal Blue
		{R: 30, G: 144, B: 255},  // Dodger Blue
		{R: 70, G: 180, B: 255},  // Sky Blue
		{R: 135, G: 206, B: 250}, // Light Sky Blue
		{R: 173, G: 216, B: 230}, // Light Blue
	}

	// Lava (Dark Red to Bright Yellow)
	Lava = []RGB{
		{R: 100, G: 0, B: 0},     // Dark Red
		{R: 180, G: 0, B: 0},     // Red
		{R: 255, G: 50, B: 0},    // Red-Orange
		{R: 255, G: 100, B: 0},   // Orange
		{R: 255, G: 180, B: 0},   // Yellow-Orange
		{R: 255, G: 255, B: 100}, // Bright Yellow
	}

	// Gruvbox Inspired
	Gruvbox = []RGB{
		{R: 251, G: 73, B: 52},   // Red
		{R: 250, G: 189, B: 47},  // Yellow
		{R: 184, G: 187, B: 38},  // Green
		{R: 142, G: 192, B: 124}, // Aqua
		{R: 131, G: 165, B: 152}, // Blue-Green
		{R: 211, G: 134, B: 155}, // Purple
	}

	// TokyoNight Tokyo Night Inspired
	TokyoNight = []RGB{
		{R: 125, G: 207, B: 255}, // Blue
		{R: 115, G: 218, B: 202}, // Cyan
		{R: 158, G: 206, B: 106}, // Green
		{R: 224, G: 175, B: 104}, // Orange
		{R: 247, G: 118, B: 142}, // Red
		{R: 187, G: 154, B: 247}, // Purple
	}
)

func CreateColoredLines(colors []RGB) string {
	var result strings.Builder

	for lineIdx, line := range lines {
		color := colors[lineIdx]
		result.WriteString(fmt.Sprintf("\x1b[38;2;%d;%d;%dm%s\x1b[0m\n",
			color.R, color.G, color.B, line))
	}

	return result.String()
}

func PrintAllSchemes() {
	schemes := map[string][]RGB{
		"Cyberpunk Neon": CyberpunkNeon,
		"Fire":           Fire,
		"Ocean":          Ocean,
		"Matrix":         Matrix,
		"Sunset":         Sunset,
		"Arctic":         Arctic,
		"Retro Amber":    RetroAmber,
		"Synthwave":      Synthwave,
		"Forest":         Forest,
		"Dracula":        Dracula,
		"Nord":           Nord,
		"Vaporwave":      Vaporwave,
		"Mono Blue":      MonoBlue,
		"Lava":           Lava,
		"Gruvbox":        Gruvbox,
		"Tokyo Night":    TokyoNight,
	}

	for name, colors := range schemes {
		fmt.Printf("\n=== %s ===\n", name)
		fmt.Println(CreateColoredLines(colors))
	}
}

func main() {
	PrintAllSchemes()
}
