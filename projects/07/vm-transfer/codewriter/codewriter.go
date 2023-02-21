package codewriter

import (
	"bytes"
	"errors"
	"fmt"
	"vm-transfer/parser"
)

type CodeWriter struct {
	filename     string
	functionName string
	namespace    string
	callIndices  map[string]int
	eqIndex      int
	gtIndex      int
	ltIndex      int
	writer       *bytes.Buffer
}

func New() *CodeWriter {
	buffer := bytes.Buffer{}

	return &CodeWriter{
		"",
		"",
		"",
		make(map[string]int),
		0,
		0,
		0,
		&buffer,
	}

}

func (c *CodeWriter) SetNamespace(namespace string) {
	c.namespace = namespace
}

func (c *CodeWriter) WritePushPop(command parser.CommandTypes, segment string, index int) error {
	code := ""

	switch command {
	case parser.PushCommand:
		code = c.handelPushCommand(segment, index)
	default:
		return errors.New("codewriter.WritePushPop only accepts PushCommand and PopCommand")
	}

	_, err := c.writer.WriteString(code)

	return err
}

func (c CodeWriter) handelPushCommand(segment string, index int) string {
	switch segment {
	case "constant":
		return fmt.Sprintf("@%d\n", index) +
			"D=A\n" +
			"@SP\n" +
			"A=M\n" +
			"M=D\n" +
			"@SP\n" +
			"M=M+1\n"
	default:
		return ""
	}
}
