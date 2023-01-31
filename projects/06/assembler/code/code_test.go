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

type compTest struct {
	in  string
	out byte
}

func TestComp(t *testing.T) {
	tests := []compTest{
		{"0", 0b0101010},
		{"1", 0b0111111},
		{"-1", 0b0111010},
		{"D", 0b0001100},
		{"A", 0b0110000},
		{"!D", 0b0001101},
		{"!A", 0b0110001},
		{"-D", 0b0001111},
		{"-A", 0b0110011},
		{"D+1", 0b0011111},
		{"A+1", 0b0110111},
		{"D-1", 0b0001110},
		{"A-1", 0b0110010},
		{"D+A", 0b0000010},
		{"D-A", 0b0010011},
		{"A-D", 0b0000111},
		{"D&A", 0b0000000},
		{"D|A", 0b0010101},
		{"M", 0b1110000},
		{"!M", 0b1110001},
		{"-M", 0b1110011},
		{"M+1", 0b1110111},
		{"M-1", 0b1110010},
		{"D+M", 0b1000010},
		{"D-M", 0b1010011},
		{"M-D", 0b1000111},
		{"D&M", 0b1000000},
		{"D|M", 0b1010101},
	}

	for i, test := range tests {
		c := New()
		comp, err := c.Comp(test.in)
		if err != nil {
			t.Errorf("#%d: error returned: %v", i, err.Error())
		}

		if comp != test.out {
			t.Errorf("#%d: got: %v want: %v", i, comp, test.out)
		}
	}
}
