package jen

// Lit renders a literal. Lit supports only built-in types (bool, string, int, complex128, float64,
// float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64).
// Passing any other type will panic.
func Lit(v interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// Lit renders a literal. Lit supports only built-in types (bool, string, int, complex128, float64,
// float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64).
// Passing any other type will panic.
func (g *Group) Lit(v interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// Lit renders a literal. Lit supports only built-in types (bool, string, int, complex128, float64,
// float32, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64).
// Passing any other type will panic.
func (s *Statement) Lit(v interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// LitFunc renders a literal. LitFunc generates the value to render by executing the provided
// function. LitFunc supports only built-in types (bool, string, int, complex128, float64, float32,
// int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64).
// Returning any other type will panic.
func LitFunc(f func() interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// LitFunc renders a literal. LitFunc generates the value to render by executing the provided
// function. LitFunc supports only built-in types (bool, string, int, complex128, float64, float32,
// int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64).
// Returning any other type will panic.
func (g *Group) LitFunc(f func() interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// LitFunc renders a literal. LitFunc generates the value to render by executing the provided
// function. LitFunc supports only built-in types (bool, string, int, complex128, float64, float32,
// int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr and complex64).
// Returning any other type will panic.
func (s *Statement) LitFunc(f func() interface{}) *Statement { _ = "STUB: not implemented"; return nil }

// LitRune renders a rune literal.
func LitRune(v rune) *Statement { _ = "STUB: not implemented"; return nil }

// LitRune renders a rune literal.
func (g *Group) LitRune(v rune) *Statement { _ = "STUB: not implemented"; return nil }

// LitRune renders a rune literal.
func (s *Statement) LitRune(v rune) *Statement { _ = "STUB: not implemented"; return nil }

// LitRuneFunc renders a rune literal. LitRuneFunc generates the value to
// render by executing the provided function.
func LitRuneFunc(f func() rune) *Statement { _ = "STUB: not implemented"; return nil }

// LitRuneFunc renders a rune literal. LitRuneFunc generates the value to
// render by executing the provided function.
func (g *Group) LitRuneFunc(f func() rune) *Statement { _ = "STUB: not implemented"; return nil }

// LitRuneFunc renders a rune literal. LitRuneFunc generates the value to
// render by executing the provided function.
func (s *Statement) LitRuneFunc(f func() rune) *Statement { _ = "STUB: not implemented"; return nil }

// LitByte renders a byte literal.
func LitByte(v byte) *Statement { _ = "STUB: not implemented"; return nil }

// LitByte renders a byte literal.
func (g *Group) LitByte(v byte) *Statement { _ = "STUB: not implemented"; return nil }

// LitByte renders a byte literal.
func (s *Statement) LitByte(v byte) *Statement { _ = "STUB: not implemented"; return nil }

// LitByteFunc renders a byte literal. LitByteFunc generates the value to
// render by executing the provided function.
func LitByteFunc(f func() byte) *Statement { _ = "STUB: not implemented"; return nil }

// LitByteFunc renders a byte literal. LitByteFunc generates the value to
// render by executing the provided function.
func (g *Group) LitByteFunc(f func() byte) *Statement { _ = "STUB: not implemented"; return nil }

// LitByteFunc renders a byte literal. LitByteFunc generates the value to
// render by executing the provided function.
func (s *Statement) LitByteFunc(f func() byte) *Statement { _ = "STUB: not implemented"; return nil }
