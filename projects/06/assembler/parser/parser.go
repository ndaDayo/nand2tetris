package parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

type Parser struct {
	currentCommand string
	reader         *bufio.Reader
	source         string
}

func New(input io.Reader) *Parser {
	b, err := ioutil.ReadAll(input)
	if err != nil {
		panic(err)
	}

	r := bytes.NewReader(b)
	return &Parser{
		"",
		bufio.NewReader(r),
		string(b),
	}
}

func (p *Parser) HasMoreCommand() bool {
	_, err := p.reader.Peek(1)
	if err != nil {
		return false
	}

	return true
}

func (p *Parser) Advance() error {
	b, _, err := p.reader.ReadLine()
	line := string(b)
	line = strings.Split(line, "//")[0]
	line = strings.Trim(line, " ")

	if err != nil {
		return err
	}
	if line == "" {
		p.Advance()
		return nil
	}

	p.currentCommand = line

	return nil
}

type CommandTypes int

const (
	A CommandTypes = iota
	C
	L
	N
	E
)

func (p *Parser) CommandType() (CommandTypes, error) {
	c := p.currentCommand
	if c == "" {
		return N, nil
	}

	if strings.HasPrefix(c, "@") {
		return A, nil
	}

	if strings.Contains(c, "=") {
		return C, nil
	}

	if strings.Contains(c, ";") {
		return C, nil
	}

	if strings.HasPrefix(c, "(") && strings.HasSuffix(c, ")") {
		return L, nil
	}

	return E, errors.New("invalid command detected")
}

func (p *Parser) Symbol() string {
	return strings.Trim(p.currentCommand, "@()")
}

func (p *Parser) Dest() string {
	if strings.Contains(p.currentCommand, "=") {
		return strings.Split(p.currentCommand, "=")[0]
	}

	return ""
}

func (p *Parser) Comp() string {
	if strings.Contains(p.currentCommand, "=") {
		return strings.Split(p.currentCommand, "=")[1]
	}

	return strings.Split(p.currentCommand, ";")[0]
}

func (p *Parser) Jump() string {
	if strings.Contains(p.currentCommand, ";") {
		return strings.Split(p.currentCommand, ";")[1]
	}

	return ""
}
