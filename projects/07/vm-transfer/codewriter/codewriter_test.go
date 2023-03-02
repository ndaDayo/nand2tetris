package codewriter

import (
	"testing"
	"vm-transfer/parser"
)

type writeArithmeticTest struct {
	command string
	out     string
}

func TestWriteArithmetic(t *testing.T) {
	tests := []writeArithmeticTest{
		{"add", "@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D+M\n@SP\nM=M+1\n"},
		{"and", "@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D&M\n@SP\nM=M+1\n"},
		{"or", "@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=D|M\n@SP\nM=M+1\n"},
		{"sub", "@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nM=M-D\n@SP\nM=M+1\n"},
		{"neg", "@SP\nM=M-1\nA=M\nM=-M\n@SP\nM=M+1\n"},
		{"not", "@SP\nM=M-1\nA=M\nM=!M\n@SP\nM=M+1\n"},
		{"eq", "@SP\nM=M-1\nA=M\nD=M\n@SP\nM=M-1\nA=M\nD=M-D\n@IF_TRUE1\nD;JEQ\n@SP\nA=M\nM=0\n@IF_END1\n0;JMP\n(IF_TRUE1)\n@SP\nA=M\nM=-1\n(IF_END1)\n@SP\nM=M+1\n"},
	}

	for i, test := range tests {
		c := New()
		c.WriteArithmetic(test.command)

		if c.writer.String() != test.out {
			t.Errorf("#%d: got: %v want: %v", i, c.writer.String(), test.out)
		}
	}
}

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
		{parser.PopCommand, "pointer", 0, "@SP\nM=M-1\nA=M\nD=M\n@THIS\nM=D\n"},
		{parser.PopCommand, "pointer", 1, "@SP\nM=M-1\nA=M\nD=M\n@THAT\nM=D\n"},
		{parser.PopCommand, "temp", 0, "@SP\nM=M-1\nA=M\nD=M\n@R5\nM=D\n"},
		{parser.PopCommand, "temp", 2, "@SP\nM=M-1\nA=M\nD=M\n@R7\nM=D\n"},
		{parser.PopCommand, "temp", 12, "@SP\nM=M-1\nA=M\nD=M\n@R17\nM=D\n"},
		{parser.PopCommand, "static", 1, "@SP\nM=M-1\nA=M\nD=M\n@Static.1\nM=D\n"},
		{parser.PopCommand, "static", 4, "@SP\nM=M-1\nA=M\nD=M\n@Static.4\nM=D\n"},
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
