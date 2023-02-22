package codewriter

import (
	"testing"
	"vm-transfer/parser"
)

type writePushPopTest struct {
	commandType parser.CommandTypes
	segment     string
	index       int
	out         string
}

func TestWritePushPop(t *testing.T) {
	tests := []writePushPopTest{
		{parser.PushCommand, "constant", 10, "@10\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "local", 8, "@LCL\nD=M\n@8\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
	}

	for i, test := range tests {
		c := New()
		c.SetNamespace("Static")
		c.WritePushPop(test.commandType, test.segment, test.index)

		if c.writer.String() != test.out {
			t.Errorf("#%d: got: %v want: %v", i, c.writer.String(), test.out)
		}
	}
}
