package ui

import (
	"errors"
	"strconv"
	"strings"
)

// AskForText requests input via stdin from the user
func (w *WriterUI) AskForText(label string) (string, error) {
	w.Systemf(label + " ")
	text, err := w.inReader.ReadString('\n')
	if err != nil {
		w.ErrorLinef("Error reading input.")
		return "", err
	}
	return strings.TrimSpace(text), nil
}

// AskForFloat requests a float value via stdin from the user
func (w *WriterUI) AskForFloat(label string) (float64, error) {
	invalid, invalidCount := true, 0
	var result float64
	for invalid && invalidCount < maxInvalidInput {
		w.Systemf(label + " ")
		text, err := w.inReader.ReadString('\n')
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
		invalidCount++
		w.ErrorLinef("Invalid float, please enter a valid value.")
	}
	if invalidCount == maxInvalidInput {
		return 0, errors.New("entered value not a valid float")
	}
	return result, nil
}

// AskForInt requests a int value via stdin from the user
func (w *WriterUI) AskForInt(label string) (int, error) {
	invalid, invalidCount := true, 0
	var result int64
	for invalid && invalidCount < maxInvalidInput {
		w.Systemf(label + " ")
		text, err := w.inReader.ReadString('\n')
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
	if invalidCount == maxInvalidInput {
		return 0, errors.New("entered value not a valid integer")
	}
	return int(result), nil
}
