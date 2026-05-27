// Package validation contains all validation functions
package validation

import (

	// x-release-please-start-major
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/config"
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/error"
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/files"

	// x-release-please-end

	"github.com/editorconfig/editorconfig-core-go/v2"
)

// keep synced with /pkg/config/config.go#L59
var textRegexes = []string{
	"^text/",
	"application/octet-stream",
	"^application/ecmascript$",
	"^application/json$",
	"^application/x-ndjson$",
	"^application/xml$",
	"\\+json",
	"\\+xml$",
}

// ValidateFile Validates a single file and returns the errors
// Note: This function is not thread safe, so it should not be called concurrently
func ValidateFile(filePath string, config config.Config) []error.ValidationError {
	_ = "STUB: not implemented"
	// idiomatic Go allows empty struct
	return nil
}

// EditorconfigConfig isn't thread safe, so we need to lock it

// ValidateFileWithDefinition Validates a single file with a given editorconfig definition and returns the errors
func ValidateFileWithDefinition(filePath string, config config.Config, def *editorconfig.Definition) []error.ValidationError {
	_ = "STUB: not implemented"
	return nil
}

// return if first line contains editorconfig-checker-disable-file

// used to ignore the line when editorconfig-checker-disable-next-line was found on previous line

// search for editorconfig-checker-enable
// but only if not disabled for performance reasons

// check for the status of the previous line (it was the next line on previous loop iteration)

// editorconfig-checker-disable-next-line was found on previous line

// check for successive editorconfig-checker-disable-next-line

// there is no need to check for editorconfig-checker-disable-line here, since line will be skipped

// skip current line

// no need to check further if disabled, for performance reasons

// a directive STARTING with editorconfig-checker-disable was found
// let's check the possible modifiers

// shorten the text for performance reasons

// this variable is here for reability, code could have been simplified, but it would have been harder to read

// check for editorconfig-checker-disable-next-line, and set status for next line

// it's not a editorconfig-checker-disable, there is no reason to disable all the following lines

// found editorconfig-checker-disable-line, skip current line

// found editorconfig-checker-disable, skip current line and all following

// ValidateFinalNewline runs the final newline validator and processes the error into the proper type
func ValidateFinalNewline(fileInformation files.FileInformation, config config.Config) error.ValidationError {
	_ = "STUB: not implemented"
	return *new(error.ValidationError)
}

// ValidateLineEnding runs the line ending validator and processes the error into the proper type
func ValidateLineEnding(fileInformation files.FileInformation, config config.Config) error.ValidationError {
	_ = "STUB: not implemented"
	return *new(error.ValidationError)
}

// ValidateIndentation runs the Indentation validator and processes the error into the proper type
func ValidateIndentation(fileInformation files.FileInformation, config config.Config) error.ValidationError {
	_ = "STUB: not implemented"
	return *new(error.ValidationError)
}

// Set indentSize to zero if there is no indentSize set

// ValidateTrailingWhitespace runs the TrailingWhitespace validator and processes the error into the proper type
func ValidateTrailingWhitespace(fileInformation files.FileInformation, config config.Config) error.ValidationError {
	_ = "STUB: not implemented"
	return *new(error.ValidationError)
}

// ValidateMaxLineLength runs the max line length validator and processes the error into the proper type
func ValidateMaxLineLength(fileInformation files.FileInformation, config config.Config) error.ValidationError {
	_ = "STUB: not implemented"
	return *new(error.ValidationError)
}

// ValidateCharset runs the charset validator and processes the error into the proper type
func ValidateCharset(fileInformation files.FileInformation, config config.Config, charset string) error.ValidationError {
	_ = "STUB: not implemented"
	return *new(error.ValidationError)
}

// ProcessValidation Validates all files and returns an array of validation errors
func ProcessValidation(files []string, config config.Config) []error.ValidationErrors {
	_ = "STUB: not implemented"
	// idiomatic Go allows empty struct
	return nil
}

// Limit the number of concurrent goroutines to the number of CPUs

// Wait for a slot to be available in the limiter, and release it via defer

// EditorconfigConfig isn't thread safe, so we need to acquire a lock

// Remove all nil values
