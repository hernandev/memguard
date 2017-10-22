package memguard

import "errors"

// ErrDestroyed is returned when a function is called on a destroyed Enclave.
var ErrDestroyed = errors.New("memguard.ErrDestroyed: buffer is destroyed")

// ErrImmutable is returned when a function that needs to modify a Enclave is given one which is immutable.
var ErrImmutable = errors.New("memguard.ErrImmutable: cannot modify immutable buffer")

// ErrInvalidLength is returned when a Enclave of smaller than one byte is requested.
var ErrInvalidLength = errors.New("memguard.ErrInvalidLength: length of buffer must be greater than zero")
