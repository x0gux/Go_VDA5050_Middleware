package task

import "errors"

var (
	ErrNotFound     = errors.New("task not found")
	ErrAlreadyExist = errors.New("task already exist")
)
