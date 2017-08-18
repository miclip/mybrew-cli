package ui

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

// AskForText requests input via stdin from the user
func (w *WriterUI) AskForText(label string) string {
	reader := bufio.NewReader(w.inReader)
	w.Systemf(label + " ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// AskForFloat requests a float value via stdin from the user
func (w *WriterUI) AskForFloat(label string) float64 {
	invalid, invalidCount := true, 0
	var result float64
	for invalid && invalidCount < MaxInvalidInput {
		reader := bufio.NewReader(w.inReader)
		w.Systemf(label + " ")
		text, err := reader.ReadString('\n')
		if err != nil {
			w.ErrorLinef("Error reading input.")
			invalidCount++
			continue
		}
		fstr := strings.TrimSpace(text)
		result, err = strconv.ParseFloat(fstr, 64)
		if err == nil {
			break
		}
		w.ErrorLinef("Invalid float, please enter a valid value.")
	}
	return result
}

// AskForInt requests a int value via stdin from the user
func (w *WriterUI) AskForInt(label string) (int, error) {
	invalid, invalidCount := true, 0
	var result int64
	for invalid && invalidCount < MaxInvalidInput {
		reader := bufio.NewReader(w.inReader)
		w.Systemf(label + " ")
		text, err := reader.ReadString('\n')
		if err != nil {
			w.ErrorLinef("Error reading input.")
			invalidCount++
			continue
		}
		fstr := strings.TrimSpace(text)
		result, err = strconv.ParseInt(fstr, 10, 0)
		if err == nil {
			break
		}
		invalidCount++
		w.ErrorLinef("Invalid int, please enter a valid value.")
	}
	if invalidCount == (MaxInvalidInput - 1) {
		return 0, errors.New("invalid value")
	}
	return int(result), nil
}
