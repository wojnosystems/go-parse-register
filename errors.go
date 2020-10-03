package parse_register

import "errors"

var ErrSettableDestination = errors.New("settableDst was not settable, please pass in a reference to the value and ensure the value is public if its in a struct")
