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

func TestWritePush(t *testing.T) {
	tests := []writePushPopTest{
		{parser.PushCommand, "constant", 10, "@10\nD=A\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "local", 8, "@LCL\nD=M\n@8\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "argument", 3, "@ARG\nD=M\n@3\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "this", 4, "@THIS\nD=M\n@4\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "that", 7, "@THAT\nD=M\n@7\nA=D+A\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "pointer", 0, "@THIS\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "pointer", 1, "@THAT\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "temp", 0, "@R5\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "temp", 2, "@R7\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "temp", 12, "@R17\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "static", 1, "@Static.1\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
		{parser.PushCommand, "static", 4, "@Static.4\nD=M\n@SP\nA=M\nM=D\n@SP\nM=M+1\n"},
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

func TestWritePop(t *testing.T) {
	tests := []writePushPopTest{
		{parser.PopCommand, "local", 1, "@SP\nM=M-1\nA=M\nD=M\n@LCL\nA=M\nA=A+1\nM=D\n"},
		{parser.PopCommand, "local", 3, "@SP\nM=M-1\nA=M\nD=M\n@LCL\nA=M\nA=A+1\nA=A+1\nA=A+1\nM=D\n"},
		{parser.PopCommand, "argument", 3, "@SP\nM=M-1\nA=M\nD=M\n@ARG\nA=M\nA=A+1\nA=A+1\nA=A+1\nM=D\n"},
		{parser.PopCommand, "this", 3, "@SP\nM=M-1\nA=M\nD=M\n@THIS\nA=M\nA=A+1\nA=A+1\nA=A+1\nM=D\n"},
		{parser.PopCommand, "that", 3, "@SP\nM=M-1\nA=M\nD=M\n@THAT\nA=M\nA=A+1\nA=A+1\nA=A+1\nM=D\n"},
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
