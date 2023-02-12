package parser

import "strings"

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

type CommandTypes int

const (
	ArithmeticCommand CommandTypes = iota
	PushCommand
	PopCommand
	LabelCommand
	GotoCommand
	IfCommand
	FunctionCommand
	ReturnCommand
	CallCommand
	UnknownCommand
)

func (p *Parser) CommandType() CommandTypes {
	commandTypeMap := map[string]CommandTypes{
		"push":     PushCommand,
		"pop":      PopCommand,
		"label":    LabelCommand,
		"goto":     GotoCommand,
		"if-goto":  IfCommand,
		"function": FunctionCommand,
		"call":     CallCommand,
		"return":   ReturnCommand,
	}

	command := p.Command()
	if cmdType, ok := commandTypeMap[command]; ok {
		return cmdType
	}

	return ArithmeticCommand
}

func (p *Parser) Command() string {
	return strings.Split(p.currentCommand, " ")[0]
}
