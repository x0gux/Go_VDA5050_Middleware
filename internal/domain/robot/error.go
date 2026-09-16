package robot

import "errors"

var (
	ErrNotFound     = errors.New("robot not found")
	ErrAlreadyExist = errors.New("robot already exist")
)
