package code

import "testing"

type destTest struct {
	in  string
	out byte
}

func TestDest(t *testing.T) {
	tests := []destTest{
		{"", 0b000},
		{"M", 0b001},
		{"D", 0b010},
		{"MD", 0b011},
		{"A", 0b100},
		{"AM", 0b101},
		{"AD", 0b110},
		{"AMD", 0b111},
	}

	for i, test := range tests {
		c := New()
		dest := c.Dest(test.in)
		if dest != test.out {
			t.Errorf("#%d: got: %v want: %v", i, dest, test.out)
		}
	}
}
