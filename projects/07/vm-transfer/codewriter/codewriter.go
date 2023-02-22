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
	segmentMap := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		"temp":     "R5",
		"constant": "",
		"static":   c.namespace,
		"pointer":  "",
	}

	segmentAddr, isSegmentMapped := segmentMap[segment]

	if !isSegmentMapped {
		return ""
	}

	switch segment {
	case "constant":
		return fmt.Sprintf("@%d\nD=A\n", index) +
			pushDToStack() +
			incrementSP()
	case "pointer":
		pointerAddr := "THIS"
		if index == 1 {
			pointerAddr = "THAT"
		}

		return fmt.Sprintf("@%s\nD=M\n", pointerAddr) +
			pushDToStack() +
			incrementSP()
	default:
		return fmt.Sprintf("@%s\nD=M\n@%d\nA=D+A\nD=M\n", segmentAddr, index) +
			pushDToStack() +
			incrementSP()
	}
}

func pushDToStack() string {
	return "@SP\nA=M\nM=D\n"
}

func incrementSP() string {
	return "@SP\nM=M+1\n"
}
