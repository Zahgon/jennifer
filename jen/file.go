package jen

// NewFile Creates a new file, with the specified package name.
func NewFile(packageName string) *File { _ = "STUB: not implemented"; return nil }

// NewFilePath creates a new file while specifying the package path - the
// package name is inferred from the path.
func NewFilePath(packagePath string) *File { _ = "STUB: not implemented"; return nil }

// NewFilePathName creates a new file with the specified package path and name.
func NewFilePathName(packagePath, packageName string) *File { _ = "STUB: not implemented"; return nil }

// File represents a single source file. Package imports are managed
// automatically by File.
type File struct {
	*Group
	name        string
	path        string
	imports     map[string]importdef
	hints       map[string]importdef
	comments    []string
	headers     []string
	cgoPreamble []string
	// NoFormat can be set to true to disable formatting of the generated source. This may be useful
	// when performance is critical, and readable code is not required.
	NoFormat bool
	// If you're worried about generated package aliases conflicting with local variable names, you
	// can set a prefix here. Package foo becomes {prefix}_foo.
	PackagePrefix string
	// CanonicalPath adds a canonical import path annotation to the package clause.
	CanonicalPath string
}

// importdef is used to differentiate packages where we know the package name from packages where the
// import is aliased. If alias == false, then name is the actual package name, and the import will be
// rendered without an alias. If used == false, the import has not been used in code yet and should be
// excluded from the import block.
type importdef struct {
	name  string
	alias bool
}

// HeaderComment adds a comment to the top of the file, above any package
// comments. A blank line is rendered below the header comments, ensuring
// header comments are not included in the package doc.
func (f *File) HeaderComment(comment string) { _ = "STUB: not implemented"; return }

// PackageComment adds a comment to the top of the file, above the package
// keyword.
func (f *File) PackageComment(comment string) { _ = "STUB: not implemented"; return }

// CgoPreamble adds a cgo preamble comment that is rendered directly before the "C" pseudo-package
// import.
func (f *File) CgoPreamble(comment string) { _ = "STUB: not implemented"; return }

// Anon adds an anonymous import.
func (f *File) Anon(paths ...string) { _ = "STUB: not implemented"; return }

// ImportName provides the package name for a path. If specified, the alias will be omitted from the
// import block. This is optional. If not specified, a sensible package name is used based on the path
// and this is added as an alias in the import block.
func (f *File) ImportName(path, name string) { _ = "STUB: not implemented"; return }

// ImportNames allows multiple names to be imported as a map. Use the [gennames](gennames) command to
// automatically generate a go file containing a map of a selection of package names.
func (f *File) ImportNames(names map[string]string) { _ = "STUB: not implemented"; return }

// ImportAlias provides the alias for a package path that should be used in the import block. A
// period can be used to force a dot-import.
func (f *File) ImportAlias(path, alias string) { _ = "STUB: not implemented"; return }

func (f *File) isLocal(path string) bool { _ = "STUB: not implemented"; return false }

func (f *File) isValidAlias(alias string) bool {
	_ = "STUB: not implemented"
	// multiple dot-imports are ok
	return false
}

// the import alias is invalid if it's a reserved word

// the import alias is invalid if it's already been registered

func (f *File) isDotImport(path string) bool { _ = "STUB: not implemented"; return false }

func (f *File) register(path string) string { _ = "STUB: not implemented"; return "" }

// notest
// should never get here because in Qual the packageToken will be null,
// so render will never be called.

// if the path has been registered previously, simply return the name

// special case for "C" pseudo-package

// look up the path in the list of provided package names and aliases by ImportName / ImportAlias

// look up the path in the list of standard library packages

// if a hint is not found for the package, guess the alias from the package path

// If the name is invalid or has been registered already, make it unique by appending a number

// If we've changed the name to make it unique, it should definitely be an alias

// Only add a prefix if the name is an alias

// Register the eventual name

// GoString renders the File for testing. Any error will cause a panic.
func (f *File) GoString() string { _ = "STUB: not implemented"; return "" }

func guessAlias(path string) string { _ = "STUB: not implemented"; return "" }

// training slashes are usually tolerated, so we can get rid of one if
// it exists

// if the path contains a "/", use the last part

// alias should be lower case

// alias should now only contain alphanumerics

// can't have a first digit, per Go identifier rules, so just skip them

// If path part was all digits, we may be left with an empty string. In this case use "pkg" as the alias.
