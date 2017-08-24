package ui

import (
	"bufio"
	"io"
	"os"

	"github.com/fatih/color"
)

var (
	maxInvalidInput = 3
)

//WriterUI Writer UI type
type WriterUI struct {
	outWriter    io.Writer
	errWriter    io.Writer
	inReader     *bufio.Reader
	errColor     *color.Color
	systemColor  *color.Color
	contentColor *color.Color
}

// NewConsoleUI creates an instance of the WriterUI
func NewConsoleUI() *WriterUI {
	return NewWriterUI(os.Stdout, os.Stderr, os.Stdin)
}

// NewWriterUI creates a UI instance
func NewWriterUI(outWriter, errWriter io.Writer, inReader io.Reader) *WriterUI {
	return &WriterUI{
		outWriter:    outWriter,
		errWriter:    errWriter,
		inReader:     bufio.NewReader(inReader),
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

// SetMaxInvalidInput overrides the max invalid input variable
func (w *WriterUI) SetMaxInvalidInput(value int) {
	maxInvalidInput = value
}

// DisplayColumns prints to stdout the items by columns
func (w *WriterUI) DisplayColumns(items []string, columns int) {
	ci := 0
	for i, v := range items {
		w.Printf("%d. %s\t", i, v)
		if (ci + 1) == columns {
			w.Printf("\n")
		}
		ci++
	}
	w.Printf("\n")
}
