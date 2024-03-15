package gropdown

import "github.com/a-h/templ"

// Position represents a position on the screen.
type Position string

func (p Position) String() string {
	return string(p)
}

// DropdownButton represents the main button for the dropdown menu.
type DropdownButton struct {
	Label string           // Label is the text displayed for the dropdown button.
	Icon  string           // Icon is the icon displayed next to the dropdown button.
	Attrs templ.Attributes // Attrs is a map of attributes to be added to the button.
}

// DropdownItem represents an item in the dropdown menu.
type DropdownItem struct {
	Label    string           // Label is the text displayed for the dropdown item.
	Icon     string           // Icon is the icon displayed next to the dropdown item.
	Href     string           // Href is the URL associated with the dropdown item (optional).
	External bool             // External if the URL associated with the dropdown item is an external URL (optional).
	Divider  bool             // Divider indicates whether the item is a divider.
	Attrs    templ.Attributes // Attrs is a map of attributes to be added to the element.
}

// Dropdown represents a dropdown menu component.
type Dropdown struct {
	Open      bool           // Open indicates whether the dropdown menu is currently open.
	Animation bool           // Animation indicates whether the dropdown button should use animations on open and close.
	Position  Position       // Position indicates the position of the dropdown content relative to the button.
	Button    DropdownButton // Button is the dropdown button configuration.
	Items     []DropdownItem // Items is a slice of dropdown menu items.
}
