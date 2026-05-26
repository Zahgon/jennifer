package jen

import (
	"io"
)

// Tag renders a struct tag
func Tag(items map[string]string) *Statement { _ = "STUB: not implemented"; return nil }

// Tag renders a struct tag
func (g *Group) Tag(items map[string]string) *Statement {
	_ = "STUB: not implemented"
	// notest
	// don't think this can ever be used in valid code?
	return nil
}

// Tag renders a struct tag
func (s *Statement) Tag(items map[string]string) *Statement { _ = "STUB: not implemented"; return nil }

type tag struct {
	items map[string]string
}

func (t tag) isNull(f *File) bool { _ = "STUB: not implemented"; return false }

func (t tag) render(f *File, w io.Writer, s *Statement) error {
	_ = "STUB: not implemented"

	// notest
	// render won't be called if t is null
	return nil
}
