package gropdown

import "github.com/a-h/templ"

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
