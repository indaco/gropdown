package gropdown

import (
	"fmt"
	"strings"
)

func buttonId(label string) string {
	return fmt.Sprintf("dropdown-button-%s", strings.ToLower(label))
}

func menuId(label string) string {
	return fmt.Sprintf("dropdown-menu-%s", strings.ToLower(label))
}

// trimmed strips away '"' from a string.
func trimmed(s string) string {
	return strings.Trim(s, "\"")
}

// toSlug returns a copy of string with lowercase
// replacing "_" and whitespaces with "-"
// example: toSlug("New Resource") returns new-resource.
func toSlug(s string) string {
	return strings.Map(func(r rune) rune {
		switch {
		case r == ' ', r == '_':
			return '-'
		case r == '"':
			return -1 // Remove quotes
		default:
			return r
		}
	}, strings.ToLower(trimmed(s)))
}

// getButtonIcon If icons are provided, use the first one; otherwise, empty ButtonIcon struct.
func getButtonIcon(icons []*ButtonIcon) *ButtonIcon {
	var icon *ButtonIcon
	if len(icons) > 0 {
		icon = icons[0]
	} else {
		icon = &ButtonIcon{}
	}
	return icon
}

// getItemOptions retrieves the options for a dropdown item, only the first one is considered.
func getItemOptions(opts []ItemOptions) ItemOptions {
	var opt ItemOptions
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = ItemOptions{}
	}
	return opt
}

// buildItemOptions creates a new DropdownItem with the given label and options.
func buildItemOptions(label string, options []ItemOptions) *DropdownItem {
	opts := getItemOptions(options)
	return &DropdownItem{
		Label:    label,
		Icon:     opts.Icon,
		Href:     opts.Href,
		External: opts.External,
		Attrs:    opts.Attrs,
	}
}
