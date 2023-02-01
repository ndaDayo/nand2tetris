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

type containsTest struct {
	in  string
	out bool
}

func TestContains(t *testing.T) {
	tests := []containsTest{
		{"contained", true},
		{"notcontained", false},
		{"sum", true},
		{"weight", false},
		{"", false},
	}

	s := New()
	s.table["contained"] = 0
	s.table["sum"] = 5

	for i, test := range tests {
		if s.Contains(test.in) != test.out {
			t.Errorf("#%d: got: %v want: %v", i, test.in, test.out)
		}
	}
}

type getAddressTest struct {
	symbol  string
	address int
}

func TestGetAddress(t *testing.T) {
	tests := []getAddressTest{
		{"i", 1},
		{"sum", 2},
		{"LOOP", 3},
		{"END", 4},
	}

	s := New()
	for _, test := range tests {
		s.AddEntry(test.symbol, test.address)
	}

	for i, test := range tests {
		address := s.GetAddress(test.symbol)
		if address != test.address {
			t.Errorf("#%d: got: %v want: %v", i, address, test.address)
		}
	}
}
