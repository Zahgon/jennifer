package jen

import (
	"io"
)

// Comment adds a comment. If the provided string contains a newline, the
// comment is formatted in multiline style. If the comment string starts
// with "//" or "/*", the automatic formatting is disabled and the string is
// rendered directly.
func Comment(str string) *Statement { _ = "STUB: not implemented"; return nil }

// Comment adds a comment. If the provided string contains a newline, the
// comment is formatted in multiline style. If the comment string starts
// with "//" or "/*", the automatic formatting is disabled and the string is
// rendered directly.
func (g *Group) Comment(str string) *Statement { _ = "STUB: not implemented"; return nil }

// Comment adds a comment. If the provided string contains a newline, the
// comment is formatted in multiline style. If the comment string starts
// with "//" or "/*", the automatic formatting is disabled and the string is
// rendered directly.
func (s *Statement) Comment(str string) *Statement { _ = "STUB: not implemented"; return nil }

// Commentf adds a comment, using a format string and a list of parameters. If
// the provided string contains a newline, the comment is formatted in
// multiline style. If the comment string starts with "//" or "/*", the
// automatic formatting is disabled and the string is rendered directly.
func Commentf(format string, a ...interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// Commentf adds a comment, using a format string and a list of parameters. If
// the provided string contains a newline, the comment is formatted in
// multiline style. If the comment string starts with "//" or "/*", the
// automatic formatting is disabled and the string is rendered directly.
func (g *Group) Commentf(format string, a ...interface{}) *Statement {
	_ = "STUB: not implemented"
	return nil
}

// Commentf adds a comment, using a format string and a list of parameters. If
// the provided string contains a newline, the comment is formatted in
// multiline style. If the comment string starts with "//" or "/*", the
// automatic formatting is disabled and the string is rendered directly.
func (s *Statement) Commentf(format string, a ...interface{}) *Statement {
	_ = "STUB: not implemented"
	return nil
}

type comment struct {
	comment string
}

func (c comment) isNull(f *File) bool { _ = "STUB: not implemented"; return false }

func (c comment) render(f *File, w io.Writer, s *Statement) error {
	_ = "STUB: not implemented"
	return nil
}

// automatic formatting disabled.
