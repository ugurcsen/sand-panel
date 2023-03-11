package domain

import "io"

// Pipes is a struct that holds the pipes of a stream.
type Pipes struct {
	Stdin  io.WriteCloser
	Stdout io.ReadCloser
	Stderr io.ReadCloser
}
