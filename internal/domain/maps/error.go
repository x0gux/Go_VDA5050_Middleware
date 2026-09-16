package maps

import "errors"

var (
	ErrNotFound         = errors.New("maps not found")
	ErrAlreadyExist     = errors.New("maps already exist")
	ErrInvalidLongitude = errors.New("longitude must be between -180 and 180")
	ErrInvalidLatitude  = errors.New("latitude must be between -90 and 90")
)
