package ui

import (
	"io"
	"os"

	"github.com/fatih/color"
)

const (
	maxInvalidInput = 3
)

//WriterUI Writer UI type
type WriterUI struct {
	outWriter    io.Writer
	errWriter    io.Writer
	errColor     *color.Color
	systemColor  *color.Color
	contentColor *color.Color
}

// NewConsoleUI creates an instance of the WriterUI
func NewConsoleUI() *WriterUI {
	return NewWriterUI(os.Stdout, os.Stderr)
}

// NewWriterUI creates a UI instance
func NewWriterUI(outWriter, errWriter io.Writer) *WriterUI {
	return &WriterUI{
		outWriter:    outWriter,
		errWriter:    errWriter,
		errColor:     color.New(color.FgRed),
		systemColor:  color.New(color.FgWhite),
		contentColor: color.New(color.FgGreen),
	}
}

// SystemLinef requests input via stdin from the user
func (w *WriterUI) SystemLinef(pattern string, args ...interface{}) {
	w.systemColor.Fprintf(w.outWriter, pattern+"\n", args...)
}

// Systemf requests input via stdin from the user
func (w *WriterUI) Systemf(pattern string, args ...interface{}) {
	w.systemColor.Fprintf(w.outWriter, pattern, args...)
}

// PrintLinef requests input via stdin from the user
func (w *WriterUI) PrintLinef(pattern string, args ...interface{}) {
	w.contentColor.Fprintf(w.outWriter, pattern+"\n", args...)
}

// Printf requests input via stdin from the user
func (w *WriterUI) Printf(pattern string, args ...interface{}) {
	w.contentColor.Fprintf(w.outWriter, pattern, args...)
}

// ErrorLinef requests input via stdin from the user
func (w *WriterUI) ErrorLinef(pattern string, args ...interface{}) {
	w.errColor.Fprintf(w.errWriter, pattern+"\n", args...)
}

// Errorf requests input via stdin from the user
func (w *WriterUI) Errorf(pattern string, args ...interface{}) {
	w.errColor.Fprintf(w.errWriter, pattern, args...)
}

// DisplayColumns prints to stdout the items by columns
func (w *WriterUI) DisplayColumns(items []string, columns int) {
	ci := 0
	for i, v := range items {
		w.contentColor.Fprintf(w.outWriter, "%d. %s\t", i, v)
		if (ci + 1) == columns {
			w.contentColor.Fprintf(w.outWriter, "\n")
		}
		ci++
	}
	w.contentColor.Fprintf(w.outWriter, "\n")
}
