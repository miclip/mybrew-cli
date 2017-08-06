package utils

import (
	"bufio"
	"math"
	"os"
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

func AskForUserInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	color.White(message)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
