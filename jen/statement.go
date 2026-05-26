package jen

import (
	"io"
)

// Statement represents a simple list of code items. When rendered the items
// are separated by spaces.
type Statement []Code

func newStatement() *Statement { _ = "STUB: not implemented"; return nil }

// Clone makes a copy of the Statement, so further tokens can be appended
// without affecting the original.
func (s *Statement) Clone() *Statement { _ = "STUB: not implemented"; return nil }

func (s *Statement) previous(c Code) Code { _ = "STUB: not implemented"; return *new(Code) }

func (s *Statement) isNull(f *File) bool { _ = "STUB: not implemented"; return false }

func (s *Statement) render(f *File, w io.Writer, _ *Statement) error {
	_ = "STUB: not implemented"
	return nil
}

// Null() token produces no output but also
// no separator. Empty() token products no
// output but adds a separator.

// Render renders the Statement to the provided writer.
func (s *Statement) Render(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// GoString renders the Statement for testing. Any error will cause a panic.
func (s *Statement) GoString() string { _ = "STUB: not implemented"; return "" }

// RenderWithFile renders the Statement to the provided writer, using imports from the provided file.
func (s *Statement) RenderWithFile(writer io.Writer, file *File) error {
	_ = "STUB: not implemented"
	return nil
}
