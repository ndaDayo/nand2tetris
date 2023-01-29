package parser

import (
	"assembler/symboltable"
	"testing"
)

var st = symboltable.New()

func TestAdvance(t *testing.T) {
	p := New("sample", st)
	p.Advance()
	if p.currentCommandIdx != 1 {
		t.Fatalf("p.currentCommandIdx should be 1, but got %d", p.currentCommandIdx)
	}
}
