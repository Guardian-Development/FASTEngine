// Package fastengine contains utility functions for working with strings.
package fastengine

import (
	"github.com/Guardian-Development/fastengine/internal/strings"
)

// Reverse returns its argument string reversed rune-wise left to right.
func Reverse(s string) string {
	return strings.InternalReverse(s)
}
