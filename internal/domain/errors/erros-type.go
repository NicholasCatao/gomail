package errortype

import "errors"

var ErrInternal error = errors.New("internal server error")
var ErrBadResquest error = errors.New("badRequest")

