package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strings"
)

func firstWordDigitInInput(line string) *int {
	var wordDigits = map[string]int{
		"one": 1, "two": 2, "three": 3,
		"four": 4, "five": 5, "six": 6,
		"seven": 7, "eight": 8, "nine": 9,
	}

	// TODO: improve brute force method
	for wordDigit, val := range wordDigits {
		if strings.Contains(line, wordDigit) {
			return &val
		}
	}

	return nil
}

func outerPositiveSingleDigits(line string, includeWordDigits bool) (firstNum, lastNum *int) {
	accumulatedInput := ""
	target := &firstNum
	for _, c := range line {
		if c >= '0' && c <= '9' {
			numberFound := int(c - '0')
			*target = &numberFound
			target = &lastNum
			accumulatedInput = ""
			continue
		}

		if !includeWordDigits {
			continue
		}

		// TODO: tidy
		accumulatedInput += string(c)
		if wordDigit := firstWordDigitInInput(accumulatedInput); wordDigit != nil {
			*target = wordDigit
			target = &lastNum
			accumulatedInput = ""
			continue
		}
	}
	if firstNum != nil && lastNum == nil {
		// Handle the case that there's only one number in the line
		lastNum = firstNum
	}
	return
}

func alignmentValueForLine(line string, includeWordDigits bool) int {
	firstNum, lastNum := outerPositiveSingleDigits(line, includeWordDigits)
	if firstNum == nil {
		return 0
	}

	const tensUnitMultiplier = 10
	slog.Debug("got vals", "first", *firstNum, "last", *lastNum)
	alignmentValue := *firstNum * tensUnitMultiplier
	alignmentValue += *lastNum
	return alignmentValue
}

func generateAlignmentValue(r io.Reader) (withoutWordDigits int, withWordDigits int) {
	reader := bufio.NewScanner(r)

	for reader.Scan() {
		line := reader.Text()
		withoutWordDigits += alignmentValueForLine(line, false)
		withWordDigits += alignmentValueForLine(line, true)
	}
	return
}

func main() {
	alignmentValuePart1, alignmentValuePart2 := generateAlignmentValue(os.Stdin)
	slog.Info("Part 1: got alignment w/o word digits", "value", alignmentValuePart1)
	slog.Info("Part 2: got alignment w/ word digits", "value", alignmentValuePart2)
}
