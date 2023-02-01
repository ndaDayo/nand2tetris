package symboltable

import "testing"

type addEntryTest struct {
	symbol  string
	address int
	total   int
}

func TestAddEntry(t *testing.T) {
	tests := []addEntryTest{
		{"i", 1, 1},
		{"sum", 2, 2},
		{"i", 1, 2},
		{"LOOP", 3, 3},
		{"sum", 2, 3},
		{"END", 4, 4},
	}

	s := New()
	for i, test := range tests {
		s.AddEntry(test.symbol, test.address)
		if len(s.table) != test.total {
			t.Errorf("#%d: got: %v want: %v", i, len(s.table), test.total)
		}
	}
}
