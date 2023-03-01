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
	segmentMap   map[string]string
}

func New() *CodeWriter {
	buffer := bytes.Buffer{}

	segmentMap := map[string]string{
		"local":    "LCL",
		"argument": "ARG",
		"this":     "THIS",
		"that":     "THAT",
		"temp":     "",
		"pointer":  "",
		"constant": "",
		"static":   "",
	}

	return &CodeWriter{
		"",
		"",
		"",
		make(map[string]int),
		0,
		0,
		0,
		&buffer,
		segmentMap,
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
	case parser.PopCommand:
		code = c.handelPopCommand(segment, index)
	default:
		return errors.New("codewriter.WritePushPop only accepts PushCommand and PopCommand")
	}

	_, err := c.writer.WriteString(code)

	return err
}

func (c CodeWriter) handelPushCommand(segment string, index int) string {
	segmentAddr, isSegmentMapped := c.segmentMap[segment]

	if !isSegmentMapped {
		return ""
	}

	switch segment {
	case "constant":
		return fmt.Sprintf("@%d\nD=A\n", index) + pushAndIncrementSP()

	case "pointer":
		pointerAddr := "THIS"
		if index == 1 {
			pointerAddr = "THAT"
		}

		return fmt.Sprintf("@%s\nD=M\n", pointerAddr) + pushAndIncrementSP()

	case "temp":
		addr := 5 + index
		return fmt.Sprintf("@R%d\nD=M\n", addr) + pushAndIncrementSP()

	case "static":
		return fmt.Sprintf("@%s.%d\nD=M\n", c.namespace, index) + pushAndIncrementSP()

	default:
		return fmt.Sprintf("@%s\nD=M\n@%d\nA=D+A\nD=M\n", segmentAddr, index) + pushAndIncrementSP()
	}
}

func pushAndIncrementSP() string {
	return "@SP\nA=M\nM=D\n@SP\nM=M+1\n"
}

func (c CodeWriter) handelPopCommand(segment string, index int) string {
	segmentAddr, isSegmentMapped := c.segmentMap[segment]

	if !isSegmentMapped {
		return ""
	}

	code := popFromStack()
	switch segment {
	case "constant":
		return ""
	default:
		code += fmt.Sprintf("@%s\nA=M\n", segmentAddr)

		for i := 0; i < index; i++ {
			code += "A=A+1\n"
		}

		code += "M=D\n"

		return code
	}
}

func popFromStack() string {
	return "@SP\nM=M-1\nA=M\nD=M\n"
}
