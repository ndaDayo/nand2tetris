package parser

import (
	"bufio"
	"bytes"
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
