package jen

import (
	"io"
)

// Dict renders as key/value pairs. Use with Values for map or composite
// literals.
type Dict map[Code]Code

// DictFunc executes a func(Dict) to generate the value. Use with Values for
// map or composite literals.
func DictFunc(f func(Dict)) Dict { _ = "STUB: not implemented"; return *new(Dict) }

func (d Dict) render(f *File, w io.Writer, s *Statement) error {
	_ = "STUB: not implemented"

	// must order keys to ensure repeatable source
	return nil
}

func (d Dict) isNull(f *File) bool { _ = "STUB: not implemented"; return false }

// if any of the key/value pairs are both not null, the Dict is not
// null
