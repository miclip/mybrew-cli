package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

// Round rounds a float to a specific precision
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// RequestUserInput requests input via stdin from the user
func RequestUserInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	color.Set(color.FgWhite)
	fmt.Print(message + " ")
	color.Unset()
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// RequestUserInputFloat requests a float value via stdin from the user
func RequestUserInputFloat(message string) float64 {
	invalid := true
	var result float64
	color.Set(color.FgWhite)
	for invalid {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(message + " ")
		text, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Error reading input.")
			continue
		}
		fstr := strings.TrimSpace(text)
		result, err = strconv.ParseFloat(fstr, 64)
		if err == nil {
			break
		}
		color.Red("Invalid float, please enter a valid value.")
	}
	color.Unset()
	return result
}

// RequestUserInputInt requests a int value via stdin from the user
func RequestUserInputInt(message string) int {
	invalid := true
	var result int64
	color.Set(color.FgWhite)
	for invalid {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(message + " ")
		text, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Error reading input.")
			continue
		}
		fstr := strings.TrimSpace(text)
		result, err = strconv.ParseInt(fstr, 10, 0)
		if err == nil {
			break
		}
		color.Red("Invalid int, please enter a valid value.")
	}
	color.Unset()
	return int(result)
}

// DisplayColumns prints to stdout the items by columns
func DisplayColumns(items []string, columns int) {
	color.Set(color.FgGreen)
	ci := 0
	for i, v := range items {
		fmt.Printf("%d. %s\t", i, v)
		if (ci + 1) == columns {
			fmt.Print("\n")
		}
		ci++
	}
	fmt.Print("\n")
	color.Unset()
}
