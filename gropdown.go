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

// DropdownBuilder is used to construct Dropdown instances with options.
type DropdownBuilder struct {
	dropdown *Dropdown
}

func (b *DropdownBuilder) Dropdown() *Dropdown {
	return b.dropdown
}

// NewDropdownBuilder creates a new DropdownBuilder instance with default settings.
func NewDropdownBuilder() *DropdownBuilder {
	return &DropdownBuilder{dropdown: &Dropdown{
		Open:      false,
		Animation: true,
		Position:  Bottom,
	}}
}

// SetOpen sets the Open field of the dropdown.
func (b *DropdownBuilder) SetOpen(open bool) *DropdownBuilder {
	b.dropdown.Open = open
	return b
}

// SetAnimation sets the animations for the dropdown button icon when open/close.
func (b *DropdownBuilder) SetAnimation(animation bool) *DropdownBuilder {
	b.dropdown.Animation = animation
	return b
}

func (b *DropdownBuilder) SetPosition(position Position) *DropdownBuilder {
	b.dropdown.Position = position
	return b
}

// WithButton sets the Button field of the dropdown.
func (b *DropdownBuilder) WithButton(button DropdownButton) *DropdownBuilder {
	b.dropdown.Button = button
	return b
}

// WithItems sets the Items field of the dropdown.
func (b *DropdownBuilder) WithItems(items []DropdownItem) *DropdownBuilder {
	b.dropdown.Items = items
	return b
}

// Render constructs and returns a templ.Component representing the dropdown.
func (b *DropdownBuilder) Render() templ.Component {
	return container(b.dropdown)
}
