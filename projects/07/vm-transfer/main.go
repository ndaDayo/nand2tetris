package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"vm-transfer/codewriter"
	"vm-transfer/parser"
)

const (
	ExitCodeOK int = iota
	ExitCodeError
)

func translateFile(path string, w *codewriter.CodeWriter) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	p := parser.New(f)

	for p.HasMoreCommands() {
		p.Advance()
		switch p.CommandType() {
		case parser.ArithmeticCommand:
			w.WriteArithmetic(p.Command())
		case parser.PushCommand, parser.PopCommand:
			index, err := p.Arg2()
			if err != nil {
				return err
			}
			w.WritePushPop(p.CommandType(), p.Arg1(), index)
		}
	}

	return nil
}

func main() {
	path := os.Args[1]
	codewriter := codewriter.New()

	extension := filepath.Ext(path)

	codewriter.SetFileName(fmt.Sprintf("%s.asm", strings.TrimSuffix(path, extension)))
	err := translateFile(path, codewriter)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(ExitCodeError)
	}

	codewriter.Save()
	os.Exit(ExitCodeOK)
}
