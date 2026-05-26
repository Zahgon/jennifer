package jen

import (
	"io"
)

// Group represents a list of Code items, separated by tokens with an optional
// open and close token.
type Group struct {
	name      string
	items     []Code
	open      string
	close     string
	separator string
	multi     bool
}

func (g *Group) isNull(f *File) bool { _ = "STUB: not implemented"; return false }

func (g *Group) isNullItems(f *File) bool { _ = "STUB: not implemented"; return false }

func (g *Group) render(f *File, w io.Writer, s *Statement) error {
	_ = "STUB: not implemented"
	return nil
}

// Special case for types - if all items are null, don't render the open/close tokens.

// Special CaseBlock format for then the previous item in the statement
// is a Case group or the default keyword.

// For multi-line blocks with a closing token, we insert a new line after the last item (but
// not if all items were null). This is to ensure that if the statement finishes with a comment,
// the closing token is not commented out.

// We also insert add trailing comma if the separator was ",".

func (g *Group) renderItems(f *File, w io.Writer) (isNull bool, err error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Special case for package tokens in Qual groups - for dot-imports, the package token
// will be null, so will not render and will not be registered in the imports block.
// This ensures all packageTokens that are rendered are registered.

// Null() token produces no output but also
// no separator. Empty() token products no
// output but adds a separator.

// The separator token is added before each non-null item, but not before the first item.

// For multi-line blocks, we insert a new line before each non-null item.

// Render renders the Group to the provided writer.
func (g *Group) Render(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// GoString renders the Group for testing. Any error will cause a panic.
func (g *Group) GoString() string { _ = "STUB: not implemented"; return "" }

// RenderWithFile renders the Group to the provided writer, using imports from the provided file.
func (g *Group) RenderWithFile(writer io.Writer, file *File) error {
	_ = "STUB: not implemented"
	return nil
}
