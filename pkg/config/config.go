// Package config contains functions and structs related to config
package config

import (
	"regexp"
	"strings"

	"github.com/editorconfig/editorconfig-core-go/v2"

	// x-release-please-start-major
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/logger"
	"github.com/editorconfig-checker/editorconfig-checker/v3/pkg/outputformat"
	// x-release-please-end
)

// DefaultExcludes is the regular expression for ignored files
var DefaultExcludes = strings.Join(defaultExcludes, "|")

// defaultExcludes are an array to produce the correct string from
var defaultExcludes = []string{
	// source control related files and folders
	"\\.git/",
	"\\.jj/",
	// package manager, generated, & lock files
	// Cargo (Rust)
	"Cargo\\.lock$",
	"/target/",
	// Composer (PHP)
	"composer\\.lock$",
	// RubyGems (Ruby)
	"Gemfile\\.lock$",
	// Go Modules (Go)
	"go\\.(mod|sum|work|work\\.sum)$",
	// Gradle (Java)
	"gradle/wrapper/gradle-wrapper\\.properties$",
	"gradlew(\\.bat)?$",
	"(buildscript-)?gradle\\.lockfile?$",
	// Maven (Java)
	"\\.mvn/wrapper/maven-wrapper\\.properties$",
	"\\.mvn/wrapper/MavenWrapperDownloader\\.java$",
	"mvnw(\\.cmd)?$",
	// NodeJS
	"/node_modules/",
	// npm (NodeJS)
	"npm-shrinkwrap\\.json$",
	"package-lock\\.json$",
	// pip (Python)
	"Pipfile\\.lock$",
	// Poetry (Python)
	"poetry\\.lock$",
	// pnpm (NodeJS)
	"pnpm-lock\\.yaml$",
	// Terraform & OpenTofu
	"\\.terraform\\.lock\\.hcl$",
	// uv (Python)
	"uv\\.lock$",
	// Buf (Protobuf)
	"buf\\.lock$",
	// yarn (NodeJS)
	"\\.pnp\\.c?js$",
	"\\.pnp\\.loader\\.mjs$",
	"\\.yarn/",
	"yarn\\.lock$",
	// font files
	"\\.eot$",
	"\\.otf$",
	"\\.ttf$",
	"\\.woff2?$",
	// image & video formats
	"\\.avif$",
	"\\.gif$",
	"\\.ico$",
	"\\.jpe?g$",
	"\\.mp4$",
	"\\.p[bgnp]m$",
	"\\.png$",
	"\\.svg$",
	"\\.tiff?$",
	"\\.webp$",
	"\\.wmv$",
	// other binary or container formats
	"\\.bak$",
	"\\.bin$",
	"\\.docx?$",
	"\\.exe$",
	"\\.pdf$",
	"\\.snap$",
	"\\.xlsx?$",
	// archive formats
	"\\.7z$",
	"\\.bz2$",
	"\\.gz$",
	"\\.jar$",
	"\\.tar$",
	"\\.tgz$",
	"\\.war$",
	"\\.zip$",
	// log & (git) patch files
	"\\.log$",
	"\\.patch$",
	// generated or minified CSS and JavaScript files
	"\\.(css|js)\\.map$",
	"min\\.(css|js)$",
	// emacs backup files
	"~$",
}

// keep synced with pkg/validation/validation.go#L20 (but no escaping)
var defaultAllowedContentTypes = []string{
	"text/",
	"application/octet-stream",
	"application/ecmascript",
	"application/json",
	"application/x-ndjson",
	"application/xml",
	"+json",
	"+xml",
}

// Config struct, contains everything a config can contain
type Config struct {
	// CLI
	ShowVersion bool
	Help        bool
	DryRun      bool
	Path        string

	// CONFIG FILE
	Version             string
	Verbose             bool
	Format              outputformat.OutputFormat
	Debug               bool
	IgnoreDefaults      bool
	SpacesAftertabs     *bool
	SpacesAfterTabs     bool
	NoColor             bool
	Exclude             []string
	AllowedContentTypes []string
	PassedFiles         []string
	Disable             DisabledChecks

	// MISC
	Logger             *logger.Logger
	EditorconfigConfig *editorconfig.Config

	// CACHE
	excludeRegexp *regexp.Regexp
}

// DisabledChecks is a Struct which represents disabled checks
type DisabledChecks struct {
	Charset                bool
	EndOfLine              bool
	Indentation            bool
	InsertFinalNewline     bool
	TrimTrailingWhitespace bool
	IndentSize             bool
	MaxLineLength          bool
}

// NewConfig initializes a new config
func NewConfig(configPaths []string) *Config { _ = "STUB: not implemented"; return nil }

// Parse parses a config at a given path
func (c *Config) Parse() error { _ = "STUB: not implemented"; return nil }

// Merge merges a provided config with a config
func (c *Config) Merge(config Config) { _ = "STUB: not implemented"; return }

// mergeDisabled merges the disabled checks into the config
// This is here because cyclomatic complexity of gocyclo was about 15 :/
func (c *Config) mergeDisabled(disabled DisabledChecks) { _ = "STUB: not implemented"; return }

// GetExcludesAsRegularExpression returns the excludes as a combined regular expression
func (c *Config) GetExcludesAsRegularExpression() string { _ = "STUB: not implemented"; return "" }

// CachedExcludesAsRegexp returns the excludes as a compiled regular expression
// The regexp compilation is cached
// Note: This is not thread-safe
func (c *Config) CachedExcludesAsRegexp() (*regexp.Regexp, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Save saves the config to it's Path
func (c Config) Save(version string) error { _ = "STUB: not implemented"; return nil }

// String returns the config in a readable form
func (c Config) String() string { _ = "STUB: not implemented"; return "" }
