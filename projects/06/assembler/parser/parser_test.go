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

type commandTypeTest struct {
	command string
	out     CommandTypes
}

func TestCommandType(t *testing.T) {
	tests := []commandTypeTest{
		{"@i", 0},
		{"@sum", 0},
		{"D=M", 1},
		{"M=M+1", 1},
		{"0;JMP", 1},
		{"(LOOP)", 2},
		{"(END)", 2},
	}

	for i, test := range tests {
		b := bytes.NewBufferString(test.command)
		p := New(b)
		p.currentCommand = test.command

		command, _ := p.CommandType()
		if command != test.out {
			t.Errorf("#%d: got: %v want: %v", i, command, test.out)
		}
	}
}

type symbolTest struct {
	in  string
	out string
}

func TestSymbol(t *testing.T) {
	tests := []symbolTest{
		{"@i", "i"},
		{"@sum", "sum"},
		{"@100", "100"},
		{"(LOOP)", "LOOP"},
		{"(END)", "END"},
	}

	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		p.currentCommand = test.in
		symbol := p.Symbol()

		if symbol != test.out {
			t.Errorf("#%d: got: %v want: %v", i, symbol, test.out)
		}
	}
}

type destTest struct {
	in  string
	out string
}

func TestDest(t *testing.T) {
	tests := []destTest{
		{"0;JMP", ""},
		{"M=M+1", "M"},
		{"D=M", "D"},
		{"MD=M-1", "MD"},
		{"A=A+1", "A"},
		{"AM=A+1", "AM"},
		{"AD=A+1", "AD"},
		{"AMD=A+1", "AMD"},
	}

	for i, test := range tests {
		b := bytes.NewBufferString(test.in)
		p := New(b)
		p.currentCommand = test.in
		dest := p.Dest()

		if dest != test.out {
			t.Errorf("#%d: got: %v want: %v", i, dest, test.out)
		}
	}

}
