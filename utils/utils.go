package utils

import (
	"bufio"
	"fmt"
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

// RequestUserInput requests input via stdin from the user
func RequestUserInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	color.Set(color.FgWhite)
	fmt.Print(message + " ")
	color.Unset()
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

// DisplayColumns prints to stdout the items by columns
func DisplayColumns(items map[int]string, columns int) {
	color.Set(color.FgGreen)

	i := 0
	for k, v := range items {
		fmt.Printf("%d. %s\t", k, v)
		if (i + 1) == columns {
			fmt.Print("\n")
		}
		i++
	}
	fmt.Print("\n")
	color.Unset()

}
