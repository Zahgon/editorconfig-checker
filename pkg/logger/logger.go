// Package logger provides functions that are logging related
package logger

import (
	"io"
	"sync"
)

// Colors which can be used
const (
	escSeqYellow = "\x1b[33;1m"
	escSeqGreen  = "\x1b[32;1m"
	escSeqRed    = "\x1b[31;1m"
	escSeqReset  = "\x1b[33;0m"
)

// Logger struct
type Logger struct {
	VerboseEnabled bool
	DebugEnabled   bool
	NoColor        bool
	writer         io.Writer
	lock           sync.Mutex
}

func GetLogger() *Logger { _ = "STUB: not implemented"; return nil }

// initialize the Logger to write to standard output
func (l *Logger) Init() { _ = "STUB: not implemented"; return }

// ensure the Logger is initialized on first print
func (l *Logger) lazyInit() { _ = "STUB: not implemented"; return }

func (l *Logger) GetWriter() io.Writer {
	_ = "STUB: not implemented"

	// allow users to overwrite the writer used
	return *new(io.Writer)
}

func (l *Logger) SetWriter(w io.Writer) {
	_ = "STUB: not implemented"

	// apply the settings from the Logger given to the instance
	return
}

func (l *Logger) Configure(newLogger *Logger) { _ = "STUB: not implemented"; return }

// Debug prints a message when Debugg is set to true on the Logger
func (l *Logger) Debug(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Verbose prints a message when Verbosee is set to true on the Logger
func (l *Logger) Verbose(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Warning prints a warning message to Stdout in yellow
func (l *Logger) Warning(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Output prints a message on Stdout in 'normal' color
func (l *Logger) Output(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Error prints an error message to Stdout in red
func (l *Logger) Error(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// println prints a message with a trailing newline
func (l *Logger) println(message string) { _ = "STUB: not implemented"; return }

// printlnColor prints a message in a given ANSI-color with a trailing newline
func (l *Logger) printlnColor(message string, color string) { _ = "STUB: not implemented"; return }
