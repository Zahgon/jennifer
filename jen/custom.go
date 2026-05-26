package jen

// Options specifies options for the Custom method
type Options struct {
	Open      string
	Close     string
	Separator string
	Multi     bool
}

// Custom renders a customized statement list. Pass in options to specify multi-line, and tokens for open, close, separator.
func Custom(options Options, statements ...Code) *Statement { _ = "STUB: not implemented"; return nil }

// Custom renders a customized statement list. Pass in options to specify multi-line, and tokens for open, close, separator.
func (g *Group) Custom(options Options, statements ...Code) *Statement {
	_ = "STUB: not implemented"
	return nil
}

// Custom renders a customized statement list. Pass in options to specify multi-line, and tokens for open, close, separator.
func (s *Statement) Custom(options Options, statements ...Code) *Statement {
	_ = "STUB: not implemented"
	return nil
}

// CustomFunc renders a customized statement list. Pass in options to specify multi-line, and tokens for open, close, separator.
func CustomFunc(options Options, f func(*Group)) *Statement { _ = "STUB: not implemented"; return nil }

// CustomFunc renders a customized statement list. Pass in options to specify multi-line, and tokens for open, close, separator.
func (g *Group) CustomFunc(options Options, f func(*Group)) *Statement {
	_ = "STUB: not implemented"
	return nil
}

// CustomFunc renders a customized statement list. Pass in options to specify multi-line, and tokens for open, close, separator.
func (s *Statement) CustomFunc(options Options, f func(*Group)) *Statement {
	_ = "STUB: not implemented"
	return nil
}
