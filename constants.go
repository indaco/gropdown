package gropdown

// ConfigContextKey is a context key for the dropdown component configurations.
var ConfigContextKey = contextKey("gropdown-config-ctx")

// Placement constants define where the dropdown content will appear on the screen.
const (
	Top    Placement = "top"
	Right  Placement = "right"
	Bottom Placement = "bottom"
	Left   Placement = "left"
)
