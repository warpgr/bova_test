package common

import "errors"

var (
	ErrLoadElement  = errors.New("can not load element")
	ErrStoreElement = errors.New("can not store element")
)
