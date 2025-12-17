// Package keyboard this documentation will be visible, when you run `go doc` cmd
package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// GetFloat reads a number from keyboard
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSuffix(input, "\n")
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
