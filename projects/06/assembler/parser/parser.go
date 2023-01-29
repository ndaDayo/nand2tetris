package parser

import (
	"assembler/symboltable"
	"assembler/value"
	"strings"
)

type Parser struct {
	*symboltable.SymbolTable
	input             string
	commandStrList    []string
	currentCommandIdx int
	readPosition      int
}

func New(input string, symbolTable *symboltable.SymbolTable) *Parser {
	parser := &Parser{
		input:             input,
		commandStrList:    strings.Split(input, value.NEW_LINE),
		currentCommandIdx: 0,
		readPosition:      0,
		SymbolTable:       symbolTable,
	}

	return parser
}

func (p *Parser) Advance() {
	p.currentCommandIdx++
}
