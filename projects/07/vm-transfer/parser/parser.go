package parser

type Parser struct {
	currentCommand string
	lines          []string
}

func (p *Parser) HasMoreCommands() bool {
	return len(p.lines) != 0
}
