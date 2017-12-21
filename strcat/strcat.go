package strcat

import "bytes"

type StringConcatenator struct {
	buffer bytes.Buffer
}

func (sc *StringConcatenator) AppendString(s string) {
	sc.buffer.WriteString(s)
}

func (sc *StringConcatenator) Clear() {
	sc.buffer = bytes.Buffer{}
}

func (sc *StringConcatenator) Len() int {
	return sc.buffer.Len()
}

func (sc *StringConcatenator) String() string {
	return sc.buffer.String()
}
