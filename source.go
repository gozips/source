package source

import "io"

// Func is the signature for the reader function that reads from the source. It
// returns the filename from the source string, a ReadCloser and an error.
// If both an ReadCloser and error sent, the entry will be created with any data
// within the ReadCloser and the error will be collected and returned to be
// handled by the endu user
type Func func(string) (string, io.ReadCloser, error)

// Readfrom calls f(s)
func (f Func) Readfrom(s string) (string, io.ReadCloser, error) {
	return f(s)
}
