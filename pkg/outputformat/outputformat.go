// Package for having structured access to our output formats
package outputformat

type OutputFormat string

const (
	Default       = OutputFormat("default")
	Codeclimate   = OutputFormat("codeclimate")
	GCC           = OutputFormat("gcc")
	GithubActions = OutputFormat("github-actions")
)

var ValidOutputFormats = []OutputFormat{
	Default,
	Codeclimate,
	GCC,
	GithubActions,
}

func GetArgumentChoiceText() string { _ = "STUB: not implemented"; return "" }

func (format OutputFormat) MarshalText() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (format *OutputFormat) UnmarshalText(data []byte) error { _ = "STUB: not implemented"; return nil }

func (format OutputFormat) IsValid() bool { _ = "STUB: not implemented"; return false }

func (f OutputFormat) String() string { _ = "STUB: not implemented"; return "" }
