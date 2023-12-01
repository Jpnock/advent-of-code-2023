package main

import (
	"io"
	"strings"
	"testing"

	"github.com/go-test/deep"
)

func newInt(val int) *int {
	return &val
}

func Test_outerNaturalNumbersInLine(t *testing.T) {
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
			gotFirstNum, gotLastNum := outerPositiveSingleDigits(tt.line)
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
	const example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`

	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Sample input",
			args: args{
				r: strings.NewReader(example),
			},
			want: 142,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateAlignmentValue(tt.args.r); got != tt.want {
				t.Errorf("generateAlignmentValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
