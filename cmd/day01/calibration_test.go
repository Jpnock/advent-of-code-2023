package main

import (
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func newInt(val int) *int {
	return &val
}

func Test_outerPositiveSingleDigits_no_word_digits(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		wantFirstNum *int
		wantLastNum  *int
	}{
		{
			name:         "exactly two - 1",
			line:         "asadsd1dasd3dasd",
			wantFirstNum: newInt(1),
			wantLastNum:  newInt(3),
		},
		{
			name:         "exactly two - 2",
			line:         "2asadsdasddasd5",
			wantFirstNum: newInt(2),
			wantLastNum:  newInt(5),
		},
		{
			name:         "one num - 1",
			line:         "asdjas3djasdh",
			wantFirstNum: newInt(3),
			wantLastNum:  newInt(3),
		},
		{
			name:         "one num - 2",
			line:         "4",
			wantFirstNum: newInt(4),
			wantLastNum:  newInt(4),
		},
		{
			name:         "no nums",
			line:         "asdjasdjasdh",
			wantFirstNum: nil,
			wantLastNum:  nil,
		},
		{
			name:         "empty line",
			line:         "",
			wantFirstNum: nil,
			wantLastNum:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstNum, gotLastNum := outerPositiveSingleDigits(tt.line, false)
			if diff := deep.Equal(gotFirstNum, tt.wantFirstNum); diff != nil {
				t.Errorf("outerNaturalNumbersInLine() gotFirstNum = %v, want %v", gotFirstNum, tt.wantFirstNum)
			}
			if diff := deep.Equal(gotLastNum, tt.wantLastNum); diff != nil {
				t.Errorf("outerNaturalNumbersInLine() gotLastNum = %v, want %v", gotLastNum, tt.wantLastNum)
			}
		})
	}
}

func Test_outerPositiveSingleDigits_with_word_digits(t *testing.T) {
	tests := []struct {
		name         string
		line         string
		wantFirstNum *int
		wantLastNum  *int
	}{
		{
			name:         "exactly two - 1",
			line:         "asadsd1dasd3dasd",
			wantFirstNum: newInt(1),
			wantLastNum:  newInt(3),
		},
		{
			name:         "exactly two - 2",
			line:         "2asadsdasddasd5",
			wantFirstNum: newInt(2),
			wantLastNum:  newInt(5),
		},
		{
			name:         "exactly two - 3",
			line:         "2asadsdasd5dasdonea",
			wantFirstNum: newInt(2),
			wantLastNum:  newInt(1),
		},
		{
			name:         "exactly two - 4",
			line:         "asadthree2sdasd5dasdonea",
			wantFirstNum: newInt(3),
			wantLastNum:  newInt(1),
		},
		{
			name:         "exactly two - 4",
			line:         "three2sdasd5dasdonea",
			wantFirstNum: newInt(3),
			wantLastNum:  newInt(1),
		},
		{
			name:         "one num - 1",
			line:         "asdjas3djasdh",
			wantFirstNum: newInt(3),
			wantLastNum:  newInt(3),
		},
		{
			name:         "one num - 2",
			line:         "4",
			wantFirstNum: newInt(4),
			wantLastNum:  newInt(4),
		},
		{
			name:         "no nums",
			line:         "asdjasdjasdh",
			wantFirstNum: nil,
			wantLastNum:  nil,
		},
		{
			name:         "empty line",
			line:         "",
			wantFirstNum: nil,
			wantLastNum:  nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFirstNum, gotLastNum := outerPositiveSingleDigits(tt.line, true)
			if diff := deep.Equal(gotFirstNum, tt.wantFirstNum); diff != nil {
				t.Errorf("outerNaturalNumbersInLine() gotFirstNum = %v, want %v", gotFirstNum, tt.wantFirstNum)
			}
			if diff := deep.Equal(gotLastNum, tt.wantLastNum); diff != nil {
				t.Errorf("outerNaturalNumbersInLine() gotLastNum = %v, want %v", gotLastNum, tt.wantLastNum)
			}
		})
	}
}

func Test_generateAlignmentValue(t *testing.T) {
	t.Run("Sample input - no word digits", func(t *testing.T) {
		const want = 142
		const example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`
		r := strings.NewReader(example)

		if got, _ := generateAlignmentValue(r); got != want {
			t.Errorf("generateAlignmentValue() = %v, want %v", got, want)
		}
	})

	t.Run("Sample input - with word digits", func(t *testing.T) {
		const want = 281
		const example = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
		r := strings.NewReader(example)
		if _, got := generateAlignmentValue(r); got != want {
			t.Errorf("generateAlignmentValue() = %v, want %v", got, want)
		}
	})
}
