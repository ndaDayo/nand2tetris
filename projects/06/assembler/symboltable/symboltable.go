package symboltable

import "fmt"

type SymbolTable struct {
	SymbolTableDist map[string]int
}

func New() *SymbolTable {
	initialSymbolTable := getInitialSymbolTable()

	return &SymbolTable{SymbolTableDist: initialSymbolTable}
}

func getInitialSymbolTable() map[string]int {
	initialSymbolTable := map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"THAT":   4,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	for i := 0; i < 16; i++ {
		initialSymbolTable[fmt.Sprintf("R%d", i)] = i
	}
	return initialSymbolTable
}
