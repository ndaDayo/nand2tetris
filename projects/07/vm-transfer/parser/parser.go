package parser

type Parser struct {
	currentCommand string
	lines          []string
}

func (p *Parser) HasMoreCommands() bool {
	return len(p.lines) != 0
}

func (p *Parser) Advance() {
	p.currentCommand = p.lines[0]
	p.lines = p.lines[1:]
}
