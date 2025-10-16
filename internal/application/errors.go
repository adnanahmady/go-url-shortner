package application

import "errors"

var (
	ErrAlreadyShorten = errors.New("url is already shorten")
	ErrUrlNotFound    = errors.New("url not found")
)
