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
