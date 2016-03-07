package errors

import (
	"errors"
	"fmt"
)

// todo: docs
func LayoutError(msg string) error {
	return fmt.Errorf("layout error: %s", msg)
}

// todo: docs
var NilModelError = errors.New("Cannot create view with nil model")
