package symboltable

type SymbolTable struct {
	table map[string]int
}

func New() *SymbolTable {
	return &SymbolTable{
		make(map[string]int),
	}
}

func (s *SymbolTable) AddEntry(symbol string, address int) {
	s.table[symbol] = address
}
