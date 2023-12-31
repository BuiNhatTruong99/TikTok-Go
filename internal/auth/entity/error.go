package entity

import "errors"

var (
	ErrUserNotFoundWithEmail = errors.New("Email is not exists")
	ErrUserNotFoundWithName  = errors.New("Name is not exists")
)
