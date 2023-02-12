package parser

import "testing"

type hasMoreCommandsTest struct {
	lines []string
	out   bool
}

func TestHasMoreCommands(t *testing.T) {
	tests := []hasMoreCommandsTest{
		{[]string{}, false},
		{[]string{"push constant 0"}, true},
		{[]string{"push constant 0", "pop local 0"}, true},
		{[]string{"push local 0", "pop pointer 1", "add"}, true},
		{[]string{"add"}, true},
	}

	for i, test := range tests {
		p := &Parser{"", test.lines}
		if p.HasMoreCommands() != test.out {
			t.Errorf("#%d: got: %v want: %v", i, p.HasMoreCommands(), test.out)
		}
	}
}

type advanceTest struct {
	before  []string
	after   []string
	command string
}

func TestAdvance(t *testing.T) {
	tests := []advanceTest{
		{[]string{"push constant 0"}, []string{}, "push constant 0"},
		{[]string{"push constant 0", "pop local 0"}, []string{"pop local 0"}, "push constant 0"},
		{[]string{"add", "sub", "eq"}, []string{"sub", "eq"}, "add"},
		{[]string{"label loop", "goto loop", "if-goto loop"}, []string{"goto loop", "if-goto loop"}, "label loop"},
		{[]string{"function Main.fibonacci 2", "return"}, []string{"return"}, "function Main.fibonacci 2"},
	}

	for i, test := range tests {
		p := &Parser{"", test.before}
		p.Advance()

		if p.currentCommand != test.command {
			t.Errorf("#%d: got: %v want: %v", i, p.currentCommand, test.command)
		}

		if len(p.lines) != len(test.after) {
			t.Errorf("#%d: got: %v want: %v", i, p.lines, test.after)
		}

		for j := range p.lines {
			if p.lines[j] != test.after[j] {
				t.Errorf("#%d: got: %v want: %v", i, p.lines[j], test.after[j])
				break
			}
		}
	}
}

type commandTypeTest struct {
	command string
	out     CommandTypes
}

func TestCommandType(t *testing.T) {
	tests := []commandTypeTest{
		{"push constant 0", PushCommand},
		{"pop location 1", PopCommand},
		{"label loop", LabelCommand},
		{"goto loop", GotoCommand},
		{"if-goto end", IfCommand},
		{"function mult 2", FunctionCommand},
		{"call mult 2 5", CallCommand},
		{"return", ReturnCommand},
		{"add", ArithmeticCommand},
		{"sub", ArithmeticCommand},
		{"lt", ArithmeticCommand},
	}

	for i, test := range tests {
		p := &Parser{test.command, []string{}}

		if p.CommandType() != test.out {
			t.Errorf("#%d: got: %v want: %v", i, p.CommandType(), test.out)
		}
	}
}

type commandTest struct {
	command string
	out     string
}

func TestCommand(t *testing.T) {
	tests := []commandTest{
		{"push constant 0", "push"},
		{"pop location 1", "pop"},
		{"label loop", "label"},
		{"goto loop", "goto"},
		{"if-goto end", "if-goto"},
		{"function mult 2", "function"},
		{"call mult 2 5", "call"},
		{"return", "return"},
		{"add", "add"},
		{"sub", "sub"},
		{"lt", "lt"},
	}

	for i, test := range tests {
		p := &Parser{test.command, []string{}}
		if p.Command() != test.out {
			t.Errorf("#%d: got: %v want: %v", i, p.Command(), test.out)
		}
	}
}
