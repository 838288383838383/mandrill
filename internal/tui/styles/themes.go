package styles

type Theme struct {
	Name      string
	BG        string
	FG        string
	Primary   string
	Secondary string
	Accent    string
	Muted     string
	Success   string
	Warning   string
	Error     string
	Border    string
	Selection string
}

var CatppuccinMocha = Theme{
	Name:      "catppuccin-mocha",
	BG:        "#1e1e2e",
	FG:        "#cdd6f4",
	Primary:   "#cba6f7",
	Secondary: "#89dceb",
	Accent:    "#f5c2e7",
	Muted:     "#6c7086",
	Success:   "#a6e3a1",
	Warning:   "#f9e2af",
	Error:     "#f38ba8",
	Border:    "#313244",
	Selection: "#45475a",
}

var CatppuccinLatte = Theme{
	Name:      "catppuccin-latte",
	BG:        "#eff1f5",
	FG:        "#4c4f69",
	Primary:   "#8839ef",
	Secondary: "#04a5e5",
	Accent:    "#ea76cb",
	Muted:     "#9ca0b0",
	Success:   "#40a02b",
	Warning:   "#df8e1d",
	Error:     "#d20f39",
	Border:    "#ccd0da",
	Selection: "#bcc0cc",
}

var CatppuccinFrappe = Theme{
	Name:      "catppuccin-frappe",
	BG:        "#303446",
	FG:        "#c6d0f5",
	Primary:   "#ca9ee6",
	Secondary: "#81c8be",
	Accent:    "#f4b8e4",
	Muted:     "#838ba7",
	Success:   "#a6d189",
	Warning:   "#e5c890",
	Error:     "#e78284",
	Border:    "#414559",
	Selection: "#51576d",
}

var Dracula = Theme{
	Name:      "dracula",
	BG:        "#282a36",
	FG:        "#f8f8f2",
	Primary:   "#bd93f9",
	Secondary: "#8be9fd",
	Accent:    "#ff79c6",
	Muted:     "#6272a4",
	Success:   "#50fa7b",
	Warning:   "#f1fa8c",
	Error:     "#ff5555",
	Border:    "#44475a",
	Selection: "#44475a",
}

var TokyoNight = Theme{
	Name:      "tokyo-night",
	BG:        "#1a1b26",
	FG:        "#c0caf5",
	Primary:   "#bb9af7",
	Secondary: "#7dcfff",
	Accent:    "#ff007c",
	Muted:     "#565f89",
	Success:   "#9ece6a",
	Warning:   "#e0af68",
	Error:     "#f7768e",
	Border:    "#24283b",
	Selection: "#33467c",
}

var Gruvbox = Theme{
	Name:      "gruvbox",
	BG:        "#282828",
	FG:        "#ebdbb2",
	Primary:   "#d3869b",
	Secondary: "#83a598",
	Accent:    "#fb4934",
	Muted:     "#928374",
	Success:   "#b8bb26",
	Warning:   "#fabd2f",
	Error:     "#fb4934",
	Border:    "#3c3836",
	Selection: "#504945",
}

var Nord = Theme{
	Name:      "nord",
	BG:        "#2e3440",
	FG:        "#d8dee9",
	Primary:   "#b48ead",
	Secondary: "#88c0d0",
	Accent:    "#bf616a",
	Muted:     "#4c566a",
	Success:   "#a3be8c",
	Warning:   "#ebcb8b",
	Error:     "#bf616a",
	Border:    "#3b4252",
	Selection: "#434c5e",
}

var BuiltInThemes = map[string]Theme{
	"catppuccin-mocha":    CatppuccinMocha,
	"catppuccin-latte":    CatppuccinLatte,
	"catppuccin-frappe":   CatppuccinFrappe,
	"dracula":             Dracula,
	"tokyo-night":         TokyoNight,
	"gruvbox":             Gruvbox,
	"nord":                Nord,
}

func GetTheme(name string) Theme {
	if t, ok := BuiltInThemes[name]; ok {
		return t
	}
	return CatppuccinMocha
}
