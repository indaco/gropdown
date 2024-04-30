package gropdown

// ConfigContextKey is a context key for the dropdown component configurations.
var ConfigContextKey = contextKey("gropdown-config-ctx")

// Position constants define where the dropdown content will appear on the screen.
const (
	Top    Position = "top"
	Right  Position = "right"
	Bottom Position = "bottom"
	Left   Position = "left"
)
