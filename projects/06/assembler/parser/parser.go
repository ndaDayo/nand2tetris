package parser

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
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
