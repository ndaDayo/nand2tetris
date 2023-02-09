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
