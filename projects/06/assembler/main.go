package main

import (
	"assembler/code"
	"assembler/parser"
	"assembler/symboltable"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
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

func (c *Client) handleSecondPass(buffer *bytes.Buffer) (bytes.Buffer, error) {
	nextAvailableRAMAddress := 16
	for c.parser.HasMoreCommand() {
		c.parser.Advance()
		commandType, err := c.parser.CommandType()

		if err != nil {
			return *buffer, err
		}

		switch commandType {
		case parser.N:
		case parser.C:
			cInstruction, err := c.handleCInstruction()
			if err != nil {
				return *buffer, err
			}

			buffer.WriteString(fmt.Sprintf("111%s\n", cInstruction))
		case parser.A:
			symbol := c.parser.Symbol()
			address, err := strconv.Atoi(symbol)
			if err != nil {
				if c.symboltable.Contains(symbol) {
					address = c.symboltable.GetAddress(symbol)
				} else {
					address = nextAvailableRAMAddress
					c.symboltable.AddEntry(symbol, address)
					nextAvailableRAMAddress++
				}
			}

			aInstruction := fmt.Sprintf("0%015b\n", address)
			buffer.WriteString(aInstruction)
		}
	}
	return *buffer, nil
}

func (c *Client) handleCInstruction() (string, error) {
	dest := c.code.Dest(c.parser.Dest())
	jump := c.code.Jump(c.parser.Jump())
	comp, err := c.code.Comp(c.parser.Comp())

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%07b%03b%03b", comp, dest, jump), nil

}
func run(r io.Reader) (bytes.Buffer, error) {
	var buffer bytes.Buffer

	client := new(r)
	err := client.handleFirstPass()
	if err != nil {
		return buffer, err
	}

	client.parser.Rewind()

	buffer, err = client.handleSecondPass(&buffer)

	if err != nil {
		return buffer, err
	}

	return buffer, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	output, err := run(reader)

	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(ExitCodeError)
	}
	fmt.Printf(output.String())
	os.Exit(ExitCodeOK)
}
