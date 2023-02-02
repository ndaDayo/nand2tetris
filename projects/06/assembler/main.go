package main

import (
	"assembler/code"
	"assembler/parser"
	"assembler/symboltable"
	"bytes"
	"fmt"
	"io"
)

type Client struct {
	parser      *parser.Parser
	code        *code.Code
	symboltable *symboltable.SymbolTable
}

func new(r io.Reader) *Client {
	s := symboltable.New()
	s.AddEntry("SP", 0)
	s.AddEntry("LCL", 1)
	s.AddEntry("ARG", 2)
	s.AddEntry("THIS", 3)
	s.AddEntry("THAT", 4)

	for i := 0; i <= 15; i++ {
		key := fmt.Sprintf("R%d", i)
		s.AddEntry(key, i)
	}

	s.AddEntry("SCREEN", 0x4000)
	s.AddEntry("KBD", 0x6000)

	return &Client{
		parser.New(r),
		code.New(),
		s,
	}
}

func (c *Client) handleFirstPass() error {
	currentAddress := 0

	for c.parser.HasMoreCommand() {
		c.parser.Advance()
		commandType, err := c.parser.CommandType()

		if err != nil {
			return err
		}
		switch commandType {
		case parser.A, parser.C:
			currentAddress++
		case parser.L:
			symbol := c.parser.Symbol()
			c.symboltable.AddEntry(symbol, currentAddress)
		}
	}

	return nil

}

func run(r io.Reader) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	client := new(r)
	err := client.handleFirstPass()
	if err != nil {
		return buffer, err
	}
}

func main() {

}
