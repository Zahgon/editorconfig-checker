// Package validators provides functions to validate if the rules of the `.editorconfig` are respected
package validators

import (
	"regexp"

	// x-release-please-start-major
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/config"
	// x-release-please-end
)

var (
	spaceRegexp              = regexp.MustCompile(`^( )*([^ \t]|$)`)
	tabRegexp                = regexp.MustCompile("^(\t)*( \\* ?|[^ \t]|$)")
	spacesAfterTabsRegexp    = regexp.MustCompile("(^(\t)*\\S)|(^(\t)+( )*\\S)|(^ \\S)")
	trailingWhitespaceRegexp = regexp.MustCompile("^.*[ \t]+$")
	finalNewlineRegexp       = regexp.MustCompile("(\n|\r|\r\n)$")
)

// Indentation validates a files indentation
func Indentation(line string, indentStyle string, indentSize int, config config.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// if no indentStyle is given it should be valid

// Space validates if a line is indented correctly respecting the indentSize
func Space(line string, indentSize int, config config.Config) error {
	_ = "STUB: not implemented"

	// match recurring spaces and everything except tab characters
	return nil
}

// match recurring spaces indentSize times - this can be recurring or never
// match either a space followed by a * and maybe a space (block-comments)
// or match everything despite a space or tab-character

// Tab validates if a line is indented with only tabs
func Tab(line string, config config.Config) error {
	_ = "STUB: not implemented"

	// match starting with one or more tabs followed by a non-whitespace char
	// OR
	// match starting with one or more tabs, followed by one space and followed by at least one non-whitespace character
	// OR
	// match starting with a space followed by at least one non-whitespace character
	return nil
}

// TrailingWhitespace validates if a line has trailing whitespace
func TrailingWhitespace(line string, trimTrailingWhitespace bool) error {
	_ = "STUB: not implemented"
	return nil
}

// FinalNewline validates if a file has a final and correct newline
func FinalNewline(fileContent string, insertFinalNewline string, endOfLine string) error {
	_ = "STUB: not implemented"
	return nil
}

// LineEnding validates if a file uses the correct line endings
func LineEnding(fileContent string, endOfLine string) error { _ = "STUB: not implemented"; return nil }

// A bit hacky because \r\n matches \r and \n

func MaxLineLength(line string, maxLineLength int, charSet string) error {
	_ = "STUB: not implemented"
	return nil
}

// strip BOM

// TODO: Handle utf-16be and utf-16le properly. Unfortunately, Go doesn't provide a utf16.RuneCountinString() function
// Just go with byte count

// Charset validates a file's charset
func Charset(charsetWanted string, charsetFound string, config config.Config) error {
	_ = "STUB: not implemented"
	return nil
}

// lowercase it as that's how they're listed in the ec spec.
