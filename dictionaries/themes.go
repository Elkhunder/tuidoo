package dictionaries

import (
	e "tuidoo/entities"

	"github.com/gdamore/tcell/v2"
)

var Themes = map[string]e.Theme{
	"dark": {
		ID:   "dark",
		Name: "Dark",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x06B6D4),
			PrimaryDark:  tcell.NewHexColor(0x0891B2),
			PrimaryLight: tcell.NewHexColor(0x22D3EE),

			Accent:      tcell.NewHexColor(0xA855F7),
			AccentDark:  tcell.NewHexColor(0x9333EA),
			AccentLight: tcell.NewHexColor(0xC084FC),

			Success: tcell.NewHexColor(0x10B981),
			Warning: tcell.NewHexColor(0xF59E0B),
			Error:   tcell.NewHexColor(0xEF4444),
			Info:    tcell.NewHexColor(0x3B82F6),

			Background: tcell.NewHexColor(0x0F172A),
			Foreground: tcell.NewHexColor(0xF1F5F9),
			Surface:    tcell.NewHexColor(0x1E293B),
			Border:     tcell.NewHexColor(0x334155),

			TextPrimary:   tcell.NewHexColor(0xF1F5F9),
			TextSecondary: tcell.NewHexColor(0x94A3B8),
			TextDisabled:  tcell.NewHexColor(0x64748B),
			TextInverse:   tcell.NewHexColor(0x0F172A),

			Hover:    tcell.NewHexColor(0x475569),
			Active:   tcell.NewHexColor(0x06B6D4),
			Focus:    tcell.NewHexColor(0x22D3EE),
			Selected: tcell.NewHexColor(0x0E7490),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x0E7490),
				tcell.NewHexColor(0x0891B2),
				tcell.NewHexColor(0x06B6D4),
				tcell.NewHexColor(0x22D3EE),
				tcell.NewHexColor(0x3B82F6),
				tcell.NewHexColor(0x0891B2),
			},
		},
	},
	"light": {
		ID:   "light",
		Name: "Light",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x0891B2),
			PrimaryDark:  tcell.NewHexColor(0x0E7490),
			PrimaryLight: tcell.NewHexColor(0x06B6D4),

			Accent:      tcell.NewHexColor(0x9333EA),
			AccentDark:  tcell.NewHexColor(0x7E22CE),
			AccentLight: tcell.NewHexColor(0xA855F7),

			Success: tcell.NewHexColor(0x059669),
			Warning: tcell.NewHexColor(0xD97706),
			Error:   tcell.NewHexColor(0xDC2626),
			Info:    tcell.NewHexColor(0x2563EB),

			Background: tcell.NewHexColor(0xF8FAFC),
			Foreground: tcell.NewHexColor(0x0F172A),
			Surface:    tcell.NewHexColor(0xFFFFFF),
			Border:     tcell.NewHexColor(0xE2E8F0),

			TextPrimary:   tcell.NewHexColor(0x0F172A),
			TextSecondary: tcell.NewHexColor(0x475569),
			TextDisabled:  tcell.NewHexColor(0x94A3B8),
			TextInverse:   tcell.NewHexColor(0xF8FAFC),

			Hover:    tcell.NewHexColor(0xF1F5F9),
			Active:   tcell.NewHexColor(0x0891B2),
			Focus:    tcell.NewHexColor(0x06B6D4),
			Selected: tcell.NewHexColor(0xCFFAFE),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x0891B2),
				tcell.NewHexColor(0x0D9488),
				tcell.NewHexColor(0x06B6D4),
				tcell.NewHexColor(0x2563EB),
				tcell.NewHexColor(0x4F46E5),
			},
		},
	},
	"mocha": {
		ID:   "mocha",
		Name: "Catppuccin Mocha",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x89B4FA),
			PrimaryDark:  tcell.NewHexColor(0x74C7EC),
			PrimaryLight: tcell.NewHexColor(0xB4BEFE),

			Accent:      tcell.NewHexColor(0xCBA6F7),
			AccentDark:  tcell.NewHexColor(0xF5C2E7),
			AccentLight: tcell.NewHexColor(0xF5E0DC),

			Success: tcell.NewHexColor(0xA6E3A1),
			Warning: tcell.NewHexColor(0xF9E2AF),
			Error:   tcell.NewHexColor(0xF38BA8),
			Info:    tcell.NewHexColor(0x89DCEB),

			Background: tcell.NewHexColor(0x1E1E2E),
			Foreground: tcell.NewHexColor(0xCDD6F4),
			Surface:    tcell.NewHexColor(0x313244),
			Border:     tcell.NewHexColor(0x45475A),

			TextPrimary:   tcell.NewHexColor(0xCDD6F4),
			TextSecondary: tcell.NewHexColor(0xBAC2DE),
			TextDisabled:  tcell.NewHexColor(0xA6ADC8),
			TextInverse:   tcell.NewHexColor(0x1E1E2E),

			Hover:    tcell.NewHexColor(0x585B70),
			Active:   tcell.NewHexColor(0x89B4FA),
			Focus:    tcell.NewHexColor(0xF5C2E7),
			Selected: tcell.NewHexColor(0x313244),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xCBA6F7),
				tcell.NewHexColor(0x89B4FA),
				tcell.NewHexColor(0x89DCEB),
				tcell.NewHexColor(0x74C7EC),
				tcell.NewHexColor(0xB4BEFE),
			},
		},
	},
	"cyberpunk": {
		ID:   "cyberpunk",
		Name: "Cyberpunk Neon",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFF00FF),
			PrimaryDark:  tcell.NewHexColor(0xC800FF),
			PrimaryLight: tcell.NewHexColor(0xFF00C8),

			Accent:      tcell.NewHexColor(0x00FFFF),
			AccentDark:  tcell.NewHexColor(0x00C8C8),
			AccentLight: tcell.NewHexColor(0x64FFFF),

			Success: tcell.NewHexColor(0x00FF96),
			Warning: tcell.NewHexColor(0xFFFF00),
			Error:   tcell.NewHexColor(0xFF0096),
			Info:    tcell.NewHexColor(0x6400FF),

			Background: tcell.NewHexColor(0x0A0014),
			Foreground: tcell.NewHexColor(0x00FFFF),
			Surface:    tcell.NewHexColor(0x1A0028),
			Border:     tcell.NewHexColor(0xFF00FF),

			TextPrimary:   tcell.NewHexColor(0x00FFFF),
			TextSecondary: tcell.NewHexColor(0xFF00C8),
			TextDisabled:  tcell.NewHexColor(0x640064),
			TextInverse:   tcell.NewHexColor(0x0A0014),

			Hover:    tcell.NewHexColor(0x320050),
			Active:   tcell.NewHexColor(0xFF00FF),
			Focus:    tcell.NewHexColor(0x00FFFF),
			Selected: tcell.NewHexColor(0xC800FF),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFF00FF),
				tcell.NewHexColor(0xFF00C8),
				tcell.NewHexColor(0xC800FF),
				tcell.NewHexColor(0x6400FF),
				tcell.NewHexColor(0x00FFFF),
			},
		},
	},
	"matrix": {
		ID:   "matrix",
		Name: "Matrix",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x00F000),
			PrimaryDark:  tcell.NewHexColor(0x00C800),
			PrimaryLight: tcell.NewHexColor(0x64FF64),

			Accent:      tcell.NewHexColor(0x00FF00),
			AccentDark:  tcell.NewHexColor(0x00C800),
			AccentLight: tcell.NewHexColor(0x96FF96),

			Success: tcell.NewHexColor(0x64FF64),
			Warning: tcell.NewHexColor(0x96FF00),
			Error:   tcell.NewHexColor(0xFF0000),
			Info:    tcell.NewHexColor(0x00A000),

			Background: tcell.NewHexColor(0x000000),
			Foreground: tcell.NewHexColor(0x00FF00),
			Surface:    tcell.NewHexColor(0x001400),
			Border:     tcell.NewHexColor(0x005000),

			TextPrimary:   tcell.NewHexColor(0x00FF00),
			TextSecondary: tcell.NewHexColor(0x00A000),
			TextDisabled:  tcell.NewHexColor(0x005000),
			TextInverse:   tcell.NewHexColor(0x000000),

			Hover:    tcell.NewHexColor(0x003200),
			Active:   tcell.NewHexColor(0x00F000),
			Focus:    tcell.NewHexColor(0x64FF64),
			Selected: tcell.NewHexColor(0x007800),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x00FF00),
				tcell.NewHexColor(0x00FF66),
				tcell.NewHexColor(0x00FFCC),
				tcell.NewHexColor(0x00CCFF),
				tcell.NewHexColor(0x0099FF),
			},
		},
	},
	"ocean": {
		ID:   "ocean",
		Name: "Ocean",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x0099CC),
			PrimaryDark:  tcell.NewHexColor(0x006699),
			PrimaryLight: tcell.NewHexColor(0x00B4DC),

			Accent:      tcell.NewHexColor(0x40E0D0),
			AccentDark:  tcell.NewHexColor(0x00B4DC),
			AccentLight: tcell.NewHexColor(0x7FFFD4),

			Success: tcell.NewHexColor(0x7FFFD4),
			Warning: tcell.NewHexColor(0xFFD700),
			Error:   tcell.NewHexColor(0xFF6B6B),
			Info:    tcell.NewHexColor(0x40E0D0),

			Background: tcell.NewHexColor(0x001928),
			Foreground: tcell.NewHexColor(0x7FFFD4),
			Surface:    tcell.NewHexColor(0x003366),
			Border:     tcell.NewHexColor(0x006699),

			TextPrimary:   tcell.NewHexColor(0x7FFFD4),
			TextSecondary: tcell.NewHexColor(0x00B4DC),
			TextDisabled:  tcell.NewHexColor(0x006699),
			TextInverse:   tcell.NewHexColor(0x001928),

			Hover:    tcell.NewHexColor(0x004C7F),
			Active:   tcell.NewHexColor(0x0099CC),
			Focus:    tcell.NewHexColor(0x40E0D0),
			Selected: tcell.NewHexColor(0x006699),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x003366),
				tcell.NewHexColor(0x006699),
				tcell.NewHexColor(0x0099CC),
				tcell.NewHexColor(0x40E0D0),
				tcell.NewHexColor(0x7FFFD4),
			},
		},
	},
	"fire": {
		ID:   "fire",
		Name: "Fire",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFF6400),
			PrimaryDark:  tcell.NewHexColor(0xFF3200),
			PrimaryLight: tcell.NewHexColor(0xFF9600),

			Accent:      tcell.NewHexColor(0xFFFF00),
			AccentDark:  tcell.NewHexColor(0xFFC800),
			AccentLight: tcell.NewHexColor(0xFFFF64),

			Success: tcell.NewHexColor(0xFFC800),
			Warning: tcell.NewHexColor(0xFF9600),
			Error:   tcell.NewHexColor(0xFF0000),
			Info:    tcell.NewHexColor(0xFF6400),

			Background: tcell.NewHexColor(0x1A0000),
			Foreground: tcell.NewHexColor(0xFFFF00),
			Surface:    tcell.NewHexColor(0x320000),
			Border:     tcell.NewHexColor(0xFF3200),

			TextPrimary:   tcell.NewHexColor(0xFFFF00),
			TextSecondary: tcell.NewHexColor(0xFF9600),
			TextDisabled:  tcell.NewHexColor(0x643200),
			TextInverse:   tcell.NewHexColor(0x1A0000),

			Hover:    tcell.NewHexColor(0x4B0000),
			Active:   tcell.NewHexColor(0xFF6400),
			Focus:    tcell.NewHexColor(0xFFFF00),
			Selected: tcell.NewHexColor(0xFF3200),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFF0000),
				tcell.NewHexColor(0xFF3200),
				tcell.NewHexColor(0xFF6400),
				tcell.NewHexColor(0xFF9600),
				tcell.NewHexColor(0xFFFF00),
			},
		},
	},
	"sunset": {
		ID:   "sunset",
		Name: "Sunset",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFF8C00),
			PrimaryDark:  tcell.NewHexColor(0xFF6432),
			PrimaryLight: tcell.NewHexColor(0xFF4564),

			Accent:      tcell.NewHexColor(0x8A2BE2),
			AccentDark:  tcell.NewHexColor(0xB400C8),
			AccentLight: tcell.NewHexColor(0xDC1496),

			Success: tcell.NewHexColor(0xFFB432),
			Warning: tcell.NewHexColor(0xFF8C00),
			Error:   tcell.NewHexColor(0xFF4564),
			Info:    tcell.NewHexColor(0x8A2BE2),

			Background: tcell.NewHexColor(0x1A0A00),
			Foreground: tcell.NewHexColor(0xFFDC96),
			Surface:    tcell.NewHexColor(0x321400),
			Border:     tcell.NewHexColor(0xFF6432),

			TextPrimary:   tcell.NewHexColor(0xFFDC96),
			TextSecondary: tcell.NewHexColor(0xFF8C00),
			TextDisabled:  tcell.NewHexColor(0x643200),
			TextInverse:   tcell.NewHexColor(0x1A0A00),

			Hover:    tcell.NewHexColor(0x4B1E00),
			Active:   tcell.NewHexColor(0xFF8C00),
			Focus:    tcell.NewHexColor(0x8A2BE2),
			Selected: tcell.NewHexColor(0xFF6432),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFF8C00),
				tcell.NewHexColor(0xFF6432),
				tcell.NewHexColor(0xFF4564),
				tcell.NewHexColor(0xDC1496),
				tcell.NewHexColor(0x8A2BE2),
			},
		},
	},
	"arctic": {
		ID:   "arctic",
		Name: "Arctic",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xA0C8EB),
			PrimaryDark:  tcell.NewHexColor(0x82B4DC),
			PrimaryLight: tcell.NewHexColor(0xBEDCF5),

			Accent:      tcell.NewHexColor(0xF0FAFF),
			AccentDark:  tcell.NewHexColor(0xDCF0FF),
			AccentLight: tcell.NewHexColor(0xFFFFFF),

			Success: tcell.NewHexColor(0xDCF0FF),
			Warning: tcell.NewHexColor(0xFFE6B4),
			Error:   tcell.NewHexColor(0xFF9696),
			Info:    tcell.NewHexColor(0xA0C8EB),

			Background: tcell.NewHexColor(0x0F1928),
			Foreground: tcell.NewHexColor(0xF0FAFF),
			Surface:    tcell.NewHexColor(0x1E2D3C),
			Border:     tcell.NewHexColor(0x6496C8),

			TextPrimary:   tcell.NewHexColor(0xF0FAFF),
			TextSecondary: tcell.NewHexColor(0xA0C8EB),
			TextDisabled:  tcell.NewHexColor(0x6496C8),
			TextInverse:   tcell.NewHexColor(0x0F1928),

			Hover:    tcell.NewHexColor(0x2D3C4B),
			Active:   tcell.NewHexColor(0xA0C8EB),
			Focus:    tcell.NewHexColor(0xF0FAFF),
			Selected: tcell.NewHexColor(0x82B4DC),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x6496C8),
				tcell.NewHexColor(0x82B4DC),
				tcell.NewHexColor(0xA0C8EB),
				tcell.NewHexColor(0xBEDCF5),
				tcell.NewHexColor(0xF0FAFF),
			},
		},
	},
	"retro-amber": {
		ID:   "retro-amber",
		Name: "Retro Amber",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFFA000),
			PrimaryDark:  tcell.NewHexColor(0xFF8200),
			PrimaryLight: tcell.NewHexColor(0xFFB432),

			Accent:      tcell.NewHexColor(0xFFDC96),
			AccentDark:  tcell.NewHexColor(0xFFC864),
			AccentLight: tcell.NewHexColor(0xFFF0D2),

			Success: tcell.NewHexColor(0xFFDC96),
			Warning: tcell.NewHexColor(0xFFB432),
			Error:   tcell.NewHexColor(0xFF6400),
			Info:    tcell.NewHexColor(0xFFA000),

			Background: tcell.NewHexColor(0x000000),
			Foreground: tcell.NewHexColor(0xFFDC96),
			Surface:    tcell.NewHexColor(0x0F0A00),
			Border:     tcell.NewHexColor(0xFF8200),

			TextPrimary:   tcell.NewHexColor(0xFFDC96),
			TextSecondary: tcell.NewHexColor(0xFFA000),
			TextDisabled:  tcell.NewHexColor(0x643200),
			TextInverse:   tcell.NewHexColor(0x000000),

			Hover:    tcell.NewHexColor(0x1E1400),
			Active:   tcell.NewHexColor(0xFFA000),
			Focus:    tcell.NewHexColor(0xFFDC96),
			Selected: tcell.NewHexColor(0xFF8200),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFF6400),
				tcell.NewHexColor(0xFF8200),
				tcell.NewHexColor(0xFFA000),
				tcell.NewHexColor(0xFFB432),
				tcell.NewHexColor(0xFFDC96),
			},
		},
	},
	"synthwave": {
		ID:   "synthwave",
		Name: "Synthwave",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFF00C8),
			PrimaryDark:  tcell.NewHexColor(0xFF0080),
			PrimaryLight: tcell.NewHexColor(0xDC00FF),

			Accent:      tcell.NewHexColor(0x00C8FF),
			AccentDark:  tcell.NewHexColor(0x6464FF),
			AccentLight: tcell.NewHexColor(0x64DCFF),

			Success: tcell.NewHexColor(0x00C8FF),
			Warning: tcell.NewHexColor(0xDC00FF),
			Error:   tcell.NewHexColor(0xFF0080),
			Info:    tcell.NewHexColor(0x6464FF),

			Background: tcell.NewHexColor(0x0A0014),
			Foreground: tcell.NewHexColor(0x00C8FF),
			Surface:    tcell.NewHexColor(0x140028),
			Border:     tcell.NewHexColor(0xFF00C8),

			TextPrimary:   tcell.NewHexColor(0x00C8FF),
			TextSecondary: tcell.NewHexColor(0xFF00C8),
			TextDisabled:  tcell.NewHexColor(0x640064),
			TextInverse:   tcell.NewHexColor(0x0A0014),

			Hover:    tcell.NewHexColor(0x28003C),
			Active:   tcell.NewHexColor(0xFF00C8),
			Focus:    tcell.NewHexColor(0x00C8FF),
			Selected: tcell.NewHexColor(0xDC00FF),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFF0080),
				tcell.NewHexColor(0xFF00C8),
				tcell.NewHexColor(0xDC00FF),
				tcell.NewHexColor(0x6464FF),
				tcell.NewHexColor(0x00C8FF),
			},
		},
	},
	"forest": {
		ID:   "forest",
		Name: "Forest",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x32B432),
			PrimaryDark:  tcell.NewHexColor(0x228B22),
			PrimaryLight: tcell.NewHexColor(0x64C832),

			Accent:      tcell.NewHexColor(0xB4F064),
			AccentDark:  tcell.NewHexColor(0x96DC32),
			AccentLight: tcell.NewHexColor(0xD2FF96),

			Success: tcell.NewHexColor(0xB4F064),
			Warning: tcell.NewHexColor(0xFFDC64),
			Error:   tcell.NewHexColor(0xFF6464),
			Info:    tcell.NewHexColor(0x64C832),

			Background: tcell.NewHexColor(0x000A00),
			Foreground: tcell.NewHexColor(0xD2FF96),
			Surface:    tcell.NewHexColor(0x001400),
			Border:     tcell.NewHexColor(0x228B22),

			TextPrimary:   tcell.NewHexColor(0xD2FF96),
			TextSecondary: tcell.NewHexColor(0x96DC32),
			TextDisabled:  tcell.NewHexColor(0x006400),
			TextInverse:   tcell.NewHexColor(0x000A00),

			Hover:    tcell.NewHexColor(0x001E00),
			Active:   tcell.NewHexColor(0x32B432),
			Focus:    tcell.NewHexColor(0xB4F064),
			Selected: tcell.NewHexColor(0x228B22),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x006400),
				tcell.NewHexColor(0x228B22),
				tcell.NewHexColor(0x32B432),
				tcell.NewHexColor(0x96DC32),
				tcell.NewHexColor(0xB4F064),
			},
		},
	},
	"dracula": {
		ID:   "dracula",
		Name: "Dracula",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xBD93F9),
			PrimaryDark:  tcell.NewHexColor(0x9370DB),
			PrimaryLight: tcell.NewHexColor(0xD0B3FF),

			Accent:      tcell.NewHexColor(0xFF79C6),
			AccentDark:  tcell.NewHexColor(0xFF55A3),
			AccentLight: tcell.NewHexColor(0xFFA0DC),

			Success: tcell.NewHexColor(0x50FA7B),
			Warning: tcell.NewHexColor(0xF1FA8C),
			Error:   tcell.NewHexColor(0xFF5555),
			Info:    tcell.NewHexColor(0x8BE9FD),

			Background: tcell.NewHexColor(0x282A36),
			Foreground: tcell.NewHexColor(0xF8F8F2),
			Surface:    tcell.NewHexColor(0x44475A),
			Border:     tcell.NewHexColor(0x6272A4),

			TextPrimary:   tcell.NewHexColor(0xF8F8F2),
			TextSecondary: tcell.NewHexColor(0xBD93F9),
			TextDisabled:  tcell.NewHexColor(0x6272A4),
			TextInverse:   tcell.NewHexColor(0x282A36),

			Hover:    tcell.NewHexColor(0x6272A4),
			Active:   tcell.NewHexColor(0xBD93F9),
			Focus:    tcell.NewHexColor(0xFF79C6),
			Selected: tcell.NewHexColor(0x44475A),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xBD93F9),
				tcell.NewHexColor(0x8BE9FD),
				tcell.NewHexColor(0x50FA7B),
				tcell.NewHexColor(0xFFB86C),
				tcell.NewHexColor(0xFF79C6),
			},
		},
	},
	"nord": {
		ID:   "nord",
		Name: "Nord",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x88C0D0),
			PrimaryDark:  tcell.NewHexColor(0x5E81AC),
			PrimaryLight: tcell.NewHexColor(0x8FBCBB),

			Accent:      tcell.NewHexColor(0x81A1C1),
			AccentDark:  tcell.NewHexColor(0x5E81AC),
			AccentLight: tcell.NewHexColor(0xA3BE8C),

			Success: tcell.NewHexColor(0xA3BE8C),
			Warning: tcell.NewHexColor(0xEBCB8B),
			Error:   tcell.NewHexColor(0xBF616A),
			Info:    tcell.NewHexColor(0x88C0D0),

			Background: tcell.NewHexColor(0x2E3440),
			Foreground: tcell.NewHexColor(0xECEFF4),
			Surface:    tcell.NewHexColor(0x3B4252),
			Border:     tcell.NewHexColor(0x4C566A),

			TextPrimary:   tcell.NewHexColor(0xECEFF4),
			TextSecondary: tcell.NewHexColor(0xD8DEE9),
			TextDisabled:  tcell.NewHexColor(0x4C566A),
			TextInverse:   tcell.NewHexColor(0x2E3440),

			Hover:    tcell.NewHexColor(0x434C5E),
			Active:   tcell.NewHexColor(0x88C0D0),
			Focus:    tcell.NewHexColor(0x8FBCBB),
			Selected: tcell.NewHexColor(0x4C566A),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x8FBCBB),
				tcell.NewHexColor(0x88C0D0),
				tcell.NewHexColor(0x81A1C1),
				tcell.NewHexColor(0x5E81AC),
				tcell.NewHexColor(0xA3BE8C),
			},
		},
	},
	"vaporwave": {
		ID:   "vaporwave",
		Name: "Vaporwave",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFF8CDC),
			PrimaryDark:  tcell.NewHexColor(0xFF71CE),
			PrimaryLight: tcell.NewHexColor(0xC896FF),

			Accent:      tcell.NewHexColor(0x64DCFF),
			AccentDark:  tcell.NewHexColor(0x78C8FF),
			AccentLight: tcell.NewHexColor(0x96F0FF),

			Success: tcell.NewHexColor(0x96F0FF),
			Warning: tcell.NewHexColor(0xFFDC96),
			Error:   tcell.NewHexColor(0xFF71CE),
			Info:    tcell.NewHexColor(0x96B4FF),

			Background: tcell.NewHexColor(0x0F0A14),
			Foreground: tcell.NewHexColor(0xF0DCFF),
			Surface:    tcell.NewHexColor(0x1E1428),
			Border:     tcell.NewHexColor(0xFF71CE),

			TextPrimary:   tcell.NewHexColor(0xF0DCFF),
			TextSecondary: tcell.NewHexColor(0xC896FF),
			TextDisabled:  tcell.NewHexColor(0x64466E),
			TextInverse:   tcell.NewHexColor(0x0F0A14),

			Hover:    tcell.NewHexColor(0x32283C),
			Active:   tcell.NewHexColor(0xFF8CDC),
			Focus:    tcell.NewHexColor(0x64DCFF),
			Selected: tcell.NewHexColor(0xC896FF),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFF71CE),
				tcell.NewHexColor(0xFF8CDC),
				tcell.NewHexColor(0xC896FF),
				tcell.NewHexColor(0x96B4FF),
				tcell.NewHexColor(0x64DCFF),
			},
		},
	},
	"mono-blue": {
		ID:   "mono-blue",
		Name: "Mono Blue",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x1E90FF),
			PrimaryDark:  tcell.NewHexColor(0x0064C8),
			PrimaryLight: tcell.NewHexColor(0x46B4FF),

			Accent:      tcell.NewHexColor(0x87CEFA),
			AccentDark:  tcell.NewHexColor(0x46B4FF),
			AccentLight: tcell.NewHexColor(0xADD8E6),

			Success: tcell.NewHexColor(0xADD8E6),
			Warning: tcell.NewHexColor(0xFFD700),
			Error:   tcell.NewHexColor(0xFF6B6B),
			Info:    tcell.NewHexColor(0x87CEFA),

			Background: tcell.NewHexColor(0x000A14),
			Foreground: tcell.NewHexColor(0xADD8E6),
			Surface:    tcell.NewHexColor(0x001428),
			Border:     tcell.NewHexColor(0x0064C8),

			TextPrimary:   tcell.NewHexColor(0xADD8E6),
			TextSecondary: tcell.NewHexColor(0x46B4FF),
			TextDisabled:  tcell.NewHexColor(0x003296),
			TextInverse:   tcell.NewHexColor(0x000A14),

			Hover:    tcell.NewHexColor(0x00283C),
			Active:   tcell.NewHexColor(0x1E90FF),
			Focus:    tcell.NewHexColor(0x87CEFA),
			Selected: tcell.NewHexColor(0x0064C8),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x003296),
				tcell.NewHexColor(0x0064C8),
				tcell.NewHexColor(0x1E90FF),
				tcell.NewHexColor(0x46B4FF),
				tcell.NewHexColor(0xADD8E6),
			},
		},
	},
	"lava": {
		ID:   "lava",
		Name: "Lava",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFF3200),
			PrimaryDark:  tcell.NewHexColor(0xB40000),
			PrimaryLight: tcell.NewHexColor(0xFF6400),

			Accent:      tcell.NewHexColor(0xFFFF64),
			AccentDark:  tcell.NewHexColor(0xFFB400),
			AccentLight: tcell.NewHexColor(0xFFFF96),

			Success: tcell.NewHexColor(0xFFB400),
			Warning: tcell.NewHexColor(0xFF6400),
			Error:   tcell.NewHexColor(0x640000),
			Info:    tcell.NewHexColor(0xFF3200),

			Background: tcell.NewHexColor(0x0A0000),
			Foreground: tcell.NewHexColor(0xFFFF96),
			Surface:    tcell.NewHexColor(0x1E0000),
			Border:     tcell.NewHexColor(0xB40000),

			TextPrimary:   tcell.NewHexColor(0xFFFF96),
			TextSecondary: tcell.NewHexColor(0xFFB400),
			TextDisabled:  tcell.NewHexColor(0x640000),
			TextInverse:   tcell.NewHexColor(0x0A0000),

			Hover:    tcell.NewHexColor(0x320000),
			Active:   tcell.NewHexColor(0xFF3200),
			Focus:    tcell.NewHexColor(0xFFFF64),
			Selected: tcell.NewHexColor(0xB40000),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x640000),
				tcell.NewHexColor(0xB40000),
				tcell.NewHexColor(0xFF3200),
				tcell.NewHexColor(0xFF6400),
				tcell.NewHexColor(0xFFFF64),
			},
		},
	},
	"gruvbox": {
		ID:   "gruvbox",
		Name: "Gruvbox",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0xFABD2F),
			PrimaryDark:  tcell.NewHexColor(0xD79921),
			PrimaryLight: tcell.NewHexColor(0xFADB4E),

			Accent:      tcell.NewHexColor(0xD3869B),
			AccentDark:  tcell.NewHexColor(0xB16286),
			AccentLight: tcell.NewHexColor(0xE5A4BA),

			Success: tcell.NewHexColor(0xB8BB26),
			Warning: tcell.NewHexColor(0xFABD2F),
			Error:   tcell.NewHexColor(0xFB4934),
			Info:    tcell.NewHexColor(0x83A598),

			Background: tcell.NewHexColor(0x282828),
			Foreground: tcell.NewHexColor(0xEBDBB2),
			Surface:    tcell.NewHexColor(0x3C3836),
			Border:     tcell.NewHexColor(0x504945),

			TextPrimary:   tcell.NewHexColor(0xEBDBB2),
			TextSecondary: tcell.NewHexColor(0xD5C4A1),
			TextDisabled:  tcell.NewHexColor(0x665C54),
			TextInverse:   tcell.NewHexColor(0x282828),

			Hover:    tcell.NewHexColor(0x504945),
			Active:   tcell.NewHexColor(0xFABD2F),
			Focus:    tcell.NewHexColor(0xD3869B),
			Selected: tcell.NewHexColor(0x504945),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0xFB4934),
				tcell.NewHexColor(0xFABD2F),
				tcell.NewHexColor(0xB8BB26),
				tcell.NewHexColor(0x83A598),
				tcell.NewHexColor(0xD3869B),
			},
		},
	},
	"tokyo-night": {
		ID:   "tokyo-night",
		Name: "Tokyo Night",
		Colors: e.Colors{
			Primary:      tcell.NewHexColor(0x7DCFFF),
			PrimaryDark:  tcell.NewHexColor(0x2AC3DE),
			PrimaryLight: tcell.NewHexColor(0xB4F9F8),

			Accent:      tcell.NewHexColor(0xBB9AF7),
			AccentDark:  tcell.NewHexColor(0x9D7CD8),
			AccentLight: tcell.NewHexColor(0xC0B6F2),

			Success: tcell.NewHexColor(0x9ECE6A),
			Warning: tcell.NewHexColor(0xE0AF68),
			Error:   tcell.NewHexColor(0xF7768E),
			Info:    tcell.NewHexColor(0x73DACA),

			Background: tcell.NewHexColor(0x1A1B26),
			Foreground: tcell.NewHexColor(0xC0CAF5),
			Surface:    tcell.NewHexColor(0x24283B),
			Border:     tcell.NewHexColor(0x414868),

			TextPrimary:   tcell.NewHexColor(0xC0CAF5),
			TextSecondary: tcell.NewHexColor(0xA9B1D6),
			TextDisabled:  tcell.NewHexColor(0x565F89),
			TextInverse:   tcell.NewHexColor(0x1A1B26),

			Hover:    tcell.NewHexColor(0x343B58),
			Active:   tcell.NewHexColor(0x7DCFFF),
			Focus:    tcell.NewHexColor(0xBB9AF7),
			Selected: tcell.NewHexColor(0x3D59A1),

			HeaderGradient: [6]tcell.Color{
				tcell.NewHexColor(0x7DCFFF),
				tcell.NewHexColor(0x73DACA),
				tcell.NewHexColor(0x9ECE6A),
				tcell.NewHexColor(0xE0AF68),
				tcell.NewHexColor(0xBB9AF7),
			},
		},
	},
}
