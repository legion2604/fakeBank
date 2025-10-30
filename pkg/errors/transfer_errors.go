package errors

import "errors"

var (
	ErrUserRecipientNotFound = errors.New("recipient not found")
	ErrInsufficientFunds     = errors.New("insufficient funds")
)
