// todo: docs
//
package timerror

import "errors"

// todo: docs
var Quit = errors.New("Quit")

// todo: docs
var TooFewFiles = errors.New("I need at least 2 files to compare")

// todo: docs
var TooManyFiles = errors.New("I can't compare more than 3 files")
