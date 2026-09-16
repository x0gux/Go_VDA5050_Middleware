package domain

import "errors"

var (
	ErrNotFound     = errors.New("not Found")
	ErrAlreadyExist = errors.New("already exist")
	ErrInvalid      = errors.New("invalid")
)
