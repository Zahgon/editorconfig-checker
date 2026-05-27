// Package error contains functions and structs related to errors
package error

import (
	// x-release-please-start-major
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/config"
	// x-release-please-end
)

// ValidationError represents one validation error
type ValidationError struct {
	LineNumber                    int
	Message                       error
	AdditionalIdenticalErrorCount int
}

// ValidationErrors represents which errors occurred in a file
type ValidationErrors struct {
	FilePath string
	Errors   []ValidationError
}

func (error1 *ValidationError) Equal(error2 ValidationError) bool {
	_ = "STUB: not implemented"
	return false
}

// GetErrorCount returns the amount of errors
func GetErrorCount(errors []ValidationErrors) int { _ = "STUB: not implemented"; return 0 }

func ConsolidateErrors(errors []ValidationError, config config.Config) []ValidationError {
	_ = "STUB: not implemented"
	return nil
}

// filter the errors, so we do not need to care about LineNumber == -1 in the loop below

// scan through the errors

// scan through the errors after the current one

// keep track of how many consecutive lines we've seen
// make sure the outer loop jumps over the consecutive errors we just found

// if they are different errors we can stop comparing messages

func PrintErrorCount(errorCount int, config config.Config) { _ = "STUB: not implemented"; return }

func PrintErrorsAsHumanReadable(errors []ValidationErrors, config config.Config) {
	_ = "STUB: not implemented"
	return
}

func PrintErrorsAsGHA(errors []ValidationErrors, config config.Config) {
	_ = "STUB: not implemented"
	return
}

// github-actions: A format dedicated for usage in Github Actions

// gcc: A format mimicking the error format from GCC.
func PrintErrorsAsGCC(errors []ValidationErrors, config config.Config) {
	_ = "STUB: not implemented"
	return
}

// codeclimate: A format that is compatible with the codeclimate format for GitLab CI.
// https://docs.gitlab.com/ee/ci/testing/code_quality.html#implement-a-custom-tool
func PrintErrorsAsCodeclimate(errors []ValidationErrors, config config.Config) {
	_ = "STUB: not implemented"
	return
}

// marshall codeclimate issues to json

// PrintErrors prints the errors to the console
func PrintErrors(errors []ValidationErrors, config config.Config) {
	_ = "STUB: not implemented"
	return
}

// codeclimate: A format that is compatible with the codeclimate format for GitLab CI.
// https://docs.gitlab.com/ee/ci/testing/code_quality.html#implement-a-custom-tool

// gcc: A format mimicking the error format from GCC.

// github-actions: A format dedicated for usage in Github Actions

// default: A human readable text format.
