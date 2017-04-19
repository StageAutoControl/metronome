package output

import (
	"fmt"
	"io"
)

// BufferOutput renders the notes as a string to the given writer (which could f.e. be stdout)
type BufferOutput struct {
	out io.Writer
}

// NewBufferOutput returns a new BufferOutput instance
func NewBufferOutput(out io.Writer) *BufferOutput {
	return &BufferOutput{out}
}

// PlayStrong plays a strong accent note
func (o *BufferOutput) PlayStrong() {
	fmt.Fprintln(o.out, "BEEP")
}

// PlayWeak plays a weak mediate note
func (o *BufferOutput) PlayWeak() {
	fmt.Fprintln(o.out, "beep")
}
