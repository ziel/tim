package errors

import "fmt"

// todo: docs
func LayoutError(msg string) error {
	return fmt.Errorf("layout error: %s", msg)
}
