package jen

import (
	"io"
)

type tokenType string

const (
	packageToken     tokenType = "package"
	identifierToken  tokenType = "identifier"
	qualifiedToken   tokenType = "qualified"
	keywordToken     tokenType = "keyword"
	operatorToken    tokenType = "operator"
	delimiterToken   tokenType = "delimiter"
	literalToken     tokenType = "literal"
	literalRuneToken tokenType = "literal_rune"
	literalByteToken tokenType = "literal_byte"
	nullToken        tokenType = "null"
	layoutToken      tokenType = "layout"
)

type token struct {
	typ     tokenType
	content interface{}
}

func (t token) isNull(f *File) bool { _ = "STUB: not implemented"; return false }

// package token is null if the path is a dot-import or the local package path

func (t token) render(f *File, w io.Writer, s *Statement) error {
	_ = "STUB: not implemented"
	return nil
}

// default constant types can be left bare

// If the formatted value is not in scientific notation, and does not have a dot, then
// we add ".0". Otherwise it will be interpreted as an int.
// See:
// https://github.com/dave/jennifer/issues/39
// https://github.com/golang/go/issues/26363

// other built-in types need specific type info

// fmt package already renders parenthesis for complex64

// Special case for Default, which must always be followed by a colon

// notest
// do nothing (should never render a null token)

// Null adds a null item. Null items render nothing and are not followed by a
// separator in lists.
func Null() *Statement { _ = "STUB: not implemented"; return nil }

// Null adds a null item. Null items render nothing and are not followed by a
// separator in lists.
func (g *Group) Null() *Statement { _ = "STUB: not implemented"; return nil }

// Null adds a null item. Null items render nothing and are not followed by a
// separator in lists.
func (s *Statement) Null() *Statement { _ = "STUB: not implemented"; return nil }

// Empty adds an empty item. Empty items render nothing but are followed by a
// separator in lists.
func Empty() *Statement { _ = "STUB: not implemented"; return nil }

// Empty adds an empty item. Empty items render nothing but are followed by a
// separator in lists.
func (g *Group) Empty() *Statement { _ = "STUB: not implemented"; return nil }

// Empty adds an empty item. Empty items render nothing but are followed by a
// separator in lists.
func (s *Statement) Empty() *Statement { _ = "STUB: not implemented"; return nil }

// Op renders the provided operator / token.
func Op(op string) *Statement { _ = "STUB: not implemented"; return nil }

// Op renders the provided operator / token.
func (g *Group) Op(op string) *Statement { _ = "STUB: not implemented"; return nil }

// Op renders the provided operator / token.
func (s *Statement) Op(op string) *Statement { _ = "STUB: not implemented"; return nil }

// Dot renders a period followed by an identifier. Use for fields and selectors.
func Dot(name string) *Statement {
	_ = "STUB: not implemented"
	// notest
	// don't think this can be used in valid code?
	return nil
}

// Dot renders a period followed by an identifier. Use for fields and selectors.
func (g *Group) Dot(name string) *Statement {
	_ = "STUB: not implemented"
	// notest
	// don't think this can be used in valid code?
	return nil
}

// Dot renders a period followed by an identifier. Use for fields and selectors.
func (s *Statement) Dot(name string) *Statement { _ = "STUB: not implemented"; return nil }

// Id renders an identifier.
func Id(name string) *Statement { _ = "STUB: not implemented"; return nil }

// Id renders an identifier.
func (g *Group) Id(name string) *Statement { _ = "STUB: not implemented"; return nil }

// Id renders an identifier.
func (s *Statement) Id(name string) *Statement { _ = "STUB: not implemented"; return nil }

// Qual renders a qualified identifier. Imports are automatically added when
// used with a File. If the path matches the local path, the package name is
// omitted. If package names conflict they are automatically renamed. Note that
// it is not possible to reliably determine the package name given an arbitrary
// package path, so a sensible name is guessed from the path and added as an
// alias. The names of all standard library packages are known so these do not
// need to be aliased. If more control is needed of the aliases, see
// [File.ImportName](#importname) or [File.ImportAlias](#importalias).
func Qual(path, name string) *Statement { _ = "STUB: not implemented"; return nil }

// Qual renders a qualified identifier. Imports are automatically added when
// used with a File. If the path matches the local path, the package name is
// omitted. If package names conflict they are automatically renamed. Note that
// it is not possible to reliably determine the package name given an arbitrary
// package path, so a sensible name is guessed from the path and added as an
// alias. The names of all standard library packages are known so these do not
// need to be aliased. If more control is needed of the aliases, see
// [File.ImportName](#importname) or [File.ImportAlias](#importalias).
func (g *Group) Qual(path, name string) *Statement { _ = "STUB: not implemented"; return nil }

// Qual renders a qualified identifier. Imports are automatically added when
// used with a File. If the path matches the local path, the package name is
// omitted. If package names conflict they are automatically renamed. Note that
// it is not possible to reliably determine the package name given an arbitrary
// package path, so a sensible name is guessed from the path and added as an
// alias. The names of all standard library packages are known so these do not
// need to be aliased. If more control is needed of the aliases, see
// [File.ImportName](#importname) or [File.ImportAlias](#importalias).
func (s *Statement) Qual(path, name string) *Statement { _ = "STUB: not implemented"; return nil }

// Line inserts a blank line.
func Line() *Statement { _ = "STUB: not implemented"; return nil }

// Line inserts a blank line.
func (g *Group) Line() *Statement { _ = "STUB: not implemented"; return nil }

// Line inserts a blank line.
func (s *Statement) Line() *Statement { _ = "STUB: not implemented"; return nil }
