package parser

import (
	"bytes"
	"testing"
)

type hasMoreCommandsTest struct {
	in  string
	out bool
}

func TestHasMoreCommand(t *testing.T) {
	tests := []hasMoreCommandsTest{
		{"", false},
		{"\n", true},
		{"@value", true},
		{"M=1", true},
		{"(LOOP)", true},
		{"D=D-A", true},
		{"0;JMP // infinite loop", true},
	}

	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		out := p.HasMoreCommand()

		if out != test.out {
			t.Errorf("#%d: input: %v, got: %v want: %v", i, test.in, out, test.out)
		}
	}
}

type advanceTest struct {
	input       string
	nextCommand string
}

func TestAdvance(t *testing.T) {
	tests := []advanceTest{
		{"@i", "@i"},
		{"@sum", "@sum"},
		{"D=M", "D=M"},
		{"\nD=A", "D=A"},
		{"// @i=0\nM=1", "M=1"},
		{"@i // comment", "@i"},
		{"M=0\n// sum=0", "M=0"},
		{"(LOOP)", "(LOOP)"},
	}

	for i, test := range tests {
		b := bytes.NewBufferString(test.input)
		p := New(b)
		err := p.Advance()

		if err != nil {
			t.Errorf("#%d: error returned: %v", i, err.Error())
		}

		if p.currentCommand != test.nextCommand {
			t.Errorf("#%d: got: %v want: %v", i, p.currentCommand, test.nextCommand)
		}
	}
}
