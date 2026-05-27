// Package files contains functions and structs related to files
package files

import (
	"io"

	"github.com/editorconfig/editorconfig-core-go/v2"

	// x-release-please-start-major
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/config"
	// x-release-please-end
)

const DefaultMimeType = "application/octet-stream"

// FileInformation is a Struct which represents some FileInformation
type FileInformation struct {
	Line         string
	Content      string
	FilePath     string
	LineNumber   int
	Editorconfig *editorconfig.Definition
}

// IsExcluded returns whether the file is excluded via arguments or config file
func IsExcluded(filePath string, config config.Config) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// AddToFiles adds a file to a slice if it isn't already in there
// and meets the requirements and returns the new slice
func AddToFiles(filePaths []string, filePath string, config config.Config) []string {
	_ = "STUB: not implemented"
	return nil
}

// hasGlobMeta reports whether the path contains any of the glob metacharacters
// recognized by filepath.Match / filepath.Glob: '*', '?', '['.
func hasGlobMeta(path string) bool { _ = "STUB: not implemented"; return false }

// resolvePassedFile expands a single --passed-file argument into one or more
// concrete paths. Paths that exist on disk are returned unchanged. Paths that
// don't exist but look like glob patterns are expanded via filepath.Glob; if
// the pattern matches nothing the argument is returned unchanged so the caller
// can surface a not-found error in the usual way.
func resolvePassedFile(passedFile string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFilesFromDirectory returns all files from a directory and its subdirectories which should be checked
func GetFilesFromDirectory(rootDir string, config config.Config) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFiles returns all files which should be checked
func GetFiles(config config.Config) ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

// Handle explicit passed files

// It is not a git repository.

// The err would be a broken symlink for example,
// so we want to program to continue but the file should not be checked

// ReadLines returns the lines from a file as a slice
func ReadLines(content string) []string { _ = "STUB: not implemented"; return nil }

// GetContentType returns the content type of a file
func GetContentType(path string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// GetContentTypeBytes returns the content type of a byte slice
func GetContentTypeBytes(fileContent io.Reader) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Always returns a valid content-type and "application/octet-stream" if no others seemed to match.

// PathExists checks whether a path of a file or directory exists or not
func PathExists(filePath string) bool { _ = "STUB: not implemented"; return false }

// GetRelativePath returns the relative path of a file from the current working directory
func GetRelativePath(filePath string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Path is already relative. No changes needed

// IsAllowedContentType returns whether the contentType is
// an allowed content type to check or not
func IsAllowedContentType(contentType string, config config.Config) bool {
	_ = "STUB: not implemented"
	/*
		why not use mimetype.EqualsAny:
		it would only match types exactly, but we allow our users to give an entire type/ category
	*/return false
}
