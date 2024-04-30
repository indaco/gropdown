package gropdown

import "github.com/a-h/templ"

// Position represents a position on the screen.
type Position string

func (p Position) String() string {
	return string(p)
}

// DropdownItem represents an item in the dropdown menu.
type DropdownItem struct {
	Label    string           // Label is the text displayed for the dropdown item.
	Icon     string           // Icon is the icon displayed next to the dropdown item.
	Href     string           // Href is the URL associated with the dropdown item (optional).
	External bool             // External if the URL associated with the dropdown item is an external URL (optional).
	Attrs    templ.Attributes // Attrs is a map of attributes to be added to the element.
}

type ButtonIcon struct {
	value string
	size  float32
}

func Icon(v string) *ButtonIcon {
	return &ButtonIcon{value: v, size: 1.25}
}
