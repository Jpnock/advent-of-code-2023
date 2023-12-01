package main

import (
	"bufio"
	"errors"
	"io"
	"log/slog"
	"os"
)

var ErrNoNumbersInLine = errors.New("the provided line contained no numbers")

func outerPositiveSingleDigits(line string) (firstNum, lastNum *int) {
	target := &firstNum
	for _, c := range line {
		if c >= '0' && c <= '9' {
			numberFound := int(c - '0')
			*target = &numberFound
			target = &lastNum
		}
	}
	if firstNum != nil && lastNum == nil {
		// Handle the case that there's only one number in the line
		lastNum = firstNum
	}
	return
}

func generateAlignmentValue(r io.Reader) int {
	reader := bufio.NewScanner(r)

	alignmentValue := 0
	for reader.Scan() {
		line := reader.Text()
		firstNum, lastNum := outerPositiveSingleDigits(line)
		if firstNum != nil {
			slog.Debug("got vals", "first", *firstNum, "last", *lastNum)
			alignmentValue += *firstNum * 10
			alignmentValue += *lastNum
		}
	}
	return alignmentValue
}

func main() {
	// Part 1
	alignmentValue := generateAlignmentValue(os.Stdin)
	slog.Info("Got alignment", "value", alignmentValue)
}
