package main

import (
	"io"

	. "github.com/dave/jennifer/jen"
)

func hints(w io.Writer, pkg, name, goListPath, filter string, standard, novendor bool) error {
	_ = "STUB: not implemented"

	// notest
	return nil
}

/*
	// <name> contains package name hints
	var <name> = map[string]string{
		...
	}
*/

func getPackages(goListPath, filter string, standard, novendor bool) (map[string]string, error) {
	_ = "STUB: not implemented"

	// notest
	return nil, nil
}

func unvendorPath(path string) string {
	_ = "STUB: not implemented"
	// notest
	return ""
}

// FindVendor looks for the last non-terminating "vendor" path element in the given import path.
// If there isn't one, FindVendor returns ok=false.
// Otherwise, FindVendor returns ok=true and the index of the "vendor".
// Copied from cmd/go/internal/load
func findVendor(path string) (index int, ok bool) {
	_ = "STUB: not implemented"
	// notest
	// Two cases, depending on internal at start of string or not.
	// The order matters: we must return the index of the final element,
	// because the final one is where the effective import path starts.
	return 0, false
}

func hasVendor(path string) bool {
	_ = "STUB: not implemented"
	// notest
	return false
}
