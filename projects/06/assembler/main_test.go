package main

import (
	"bytes"
	"fmt"
	"testing"
)

type runTest struct {
	in  string
	out string
}

func TestRun(t *testing.T) {
	tests := []runTest{
		{"@100", fmt.Sprintf("0%015b\n", 100)},
		{"D=M", fmt.Sprintf("%b\n", 0b1111110000010000)},
		{"D;JLE", fmt.Sprintf("%b\n", 0b1110001100000110)},
		{"D=A", fmt.Sprintf("%b\n", 0b1110110000010000)},
		{"D=D+A", fmt.Sprintf("%b\n", 0b1110000010010000)},
		{"M=D", fmt.Sprintf("%b\n", 0b1110001100001000)},
		{"MD=M-1", fmt.Sprintf("%b\n", 0b1111110010011000)},
		{"D;JGT", fmt.Sprintf("%b\n", 0b1110001100000001)},
		{"// comment", ""},
		{"@100 // comment", fmt.Sprintf("0%015b\n", 100)},
		{"(LOOP)", ""},
		{"(END)\n@END\n0;JMP", fmt.Sprintf("0%015b\n%b\n", 0, 0b1110101010000111)},
		{"@i\nM=D", fmt.Sprintf("0%015b\n%b\n", 16, 0b1110001100001000)},
	}

	for i, test := range tests {
		b, err := run(bytes.NewBufferString(test.in))
		if err != nil {
			t.Errorf("#%d: error occurred: %v", i, err.Error())
		}
		if b.String != test.out {
			t.Errorf("#%d: got: %v want: %v", i, b.String(), test.out)
		}
	}
}
