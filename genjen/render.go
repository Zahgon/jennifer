package main

import (
	"io"

	. "github.com/dave/jennifer/jen"
)

func render(w io.Writer) error { _ = "STUB: not implemented"; return nil }

// b used in closures

/*
	// <comment>
	func (s *Statement) <name>(<funcParams>) *Statement {
		g := &Group{
			items:     []Code{<paramNames>}|<paramNames[0]>,
			name:      "<name>",
			open:      "<opening>",
			close:     "<closing>",
			separator: "<separator>",
			multi:     <multi>,
		}
		*s = append(*s, g)
		return s
	}
*/

/*
	// <funcComment>
	func (s *Statement) <funcName>(f func(*Group)) *Statement {
		g := &Group{
			name:      "<name>",
			open:      "<opening>",
			close:     "<closing>",
			separator: "<separator>",
			multi:     <multi>,
		}
		f(g)
		*s = append(*s, g)
		return s
	}
*/

// used in closures

// only enforce test coverage on one item

/*
	// <comment>
	func (s *Statement) <name>() *Statement {
		t := token{
			typ:     <tokenType>,
			content: "<token>",
		}
		*s = append(*s, t)
		return s
	}
*/

// only enforce test coverage on one item

// For each method on *Statement, this generates a package level
// function and a method on *Group, both with the same name.
func addFunctionAndGroupMethod(
	file *File,
	name string,
	comment *Statement,
	funcParams []Code,
	callParams []Code,
	notest bool,
) {
	_ = "STUB: not implemented"
	/*
		// <comment>
		func <name>(<funcParams>) *Statement {
			return newStatement().<name>(<callParams>)
		}
	*/return
}

// only enforce test coverage on one item

/*
	// <comment>
	func (g *Group) <name>(<funcParams>) *Statement {
		s := <name>(<callParams>)
		g.items = append(g.items, s)
		return s
	}
*/

// only enforce test coverage on one item
