package parse

import "errors"

var (
	ErrExtractPEM   = errors.New("cannot extract PEM data")
	ErrExtractBlock = errors.New("cannot extract relevant block")
)
