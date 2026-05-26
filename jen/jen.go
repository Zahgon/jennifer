// Package jen is a code generator for Go
package jen

import (
	"io"
)

// Code represents an item of code that can be rendered.
type Code interface {
	render(f *File, w io.Writer, s *Statement) error
	isNull(f *File) bool
}

// Save renders the file and saves to the filename provided.
func (f *File) Save(filename string) error {
	_ = "STUB: not implemented"
	// notest
	return nil
}

// Render renders the file to the provided writer.
func (f *File) Render(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// Append an extra newline so that header comments don't get lumped in
// with package comments.

func (f *File) renderImports(source io.Writer) error {
	_ = "STUB: not implemented"

	// Render the "C" import if it's been used in a `Qual`, `Anon` or if there's a preamble comment
	return nil
}

// Only separate the import from the main imports block if there's a preamble

// filter out the "C" pseudo-package so it's not rendered in a block with the other
// imports, but only if it is accompanied by a preamble comment

// "C" package should be rendered without alias even when used as an anonymous import
// (e.g. should never have an underscore).

// We must sort the imports to ensure repeatable
// source.

// "C" package should be rendered without alias even when used as an anonymous import
// (e.g. should never have an underscore).
