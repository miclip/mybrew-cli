package fakes

import (
	"io"

	"github.com/miclip/mybrewgo/ui"
)

// NewFakeUI returns a fake UI
func NewFakeUI(outWriter, errWriter io.Writer) *ui.WriterUI {
	return ui.NewWriterUI(outWriter, errWriter)
}
