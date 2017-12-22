// Package strcat provides a simple wrapper around the bytes.Buffer struct to implement simple string builder functionality.
package strcat

import "bytes"

// StringConcatenator implements simple string builder functionality.
type StringConcatenator struct {
	buffer bytes.Buffer
}

// AppendString copies the supplied string to the end of the buffer.
func (sc *StringConcatenator) AppendString(s string) {
	sc.buffer.WriteString(s)
}

// Clear empties the buffer.
func (sc *StringConcatenator) Clear() {
	sc.buffer = bytes.Buffer{}
}

// Len returns the length of the buffer.
func (sc *StringConcatenator) Len() int {
	return sc.buffer.Len()
}

func (sc *StringConcatenator) String() string {
	return sc.buffer.String()
}
