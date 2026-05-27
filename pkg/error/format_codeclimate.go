package error

// CodeclimageLines represents the lines of an issue in codeclimate format
type CodeclimateLines struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

// CodeclimateLocation represents the location of an issue in codeclimate format
type CodeclimateLocation struct {
	Path  string           `json:"path"`
	Lines CodeclimateLines `json:"lines"`
}

// CodeclimateIssue represents an issue in codeclimate format
type CodeclimateIssue struct {
	Check       string              `json:"check_name"`
	Description string              `json:"description"`
	Fingerprint string              `json:"fingerprint"`
	Severity    string              `json:"severity"`
	Location    CodeclimateLocation `json:"location"`
}

const (
	checkName = "editorconfig-checker"
	severity  = "minor"
)

func newCodeclimateIssue(err ValidationError, path string) CodeclimateIssue {
	_ = "STUB: not implemented"
	return *new(CodeclimateIssue)
}
