package codewriter

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"vm-transfer/parser"
)

type CodeWriter struct {
	filename     string
	functionName string
	namespace    string
	callIndices  map[string]int
	labelId      int
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
		&buffer,
		segmentMap,
	}

}

func (c *CodeWriter) SetFileName(filename string) {
	c.filename = filename
}

func (c *CodeWriter) SetNamespace(namespace string) {
	c.namespace = namespace
}

func (c *CodeWriter) WriteArithmetic(command string) error {
	code := ""

	switch command {
	case "add", "and", "or":
		code = popFromStack() + "D=M\n" + popFromStack()
		op, err := binaryCommandOperator(command)

		if err != nil {
			return err
		}

		code += fmt.Sprintf("M=D%sM\n", op) + incrementSP()
	case "sub":
		code = popFromStack() + "D=M\n" + popFromStack()
		code += fmt.Sprintf("M=M-D\n") + incrementSP()
	case "neg", "not":
		op, err := unaryCommandOperator(command)
		if err != nil {
			return err
		}
		code = popFromStack() + fmt.Sprintf("M=%sM\n", op) + incrementSP()

	case "eq", "gt", "lt":
		c.labelId++

		code = popFromStack() + "D=M\n" + popFromStack() + "D=M-D\n"

		trueLabel := fmt.Sprintf("IF_TRUE%d", c.labelId)
		endLabel := fmt.Sprintf("IF_END%d", c.labelId)

		jumpMap := map[string]string{
			"eq": "JEQ",
			"gt": "JGT",
			"lt": "JLT",
		}

		jump, ok := jumpMap[command]
		if !ok {
			return fmt.Errorf("invalid command %s", command)
		}

		code += fmt.Sprintf("@%s\nD;%s\n", trueLabel, jump)
		code += pushToStack("0")
		code += fmt.Sprintf("@%s\n0;JMP\n", endLabel)
		code += fmt.Sprintf("(%s)\n", trueLabel)
		code += pushToStack("-1")
		code += fmt.Sprintf("(%s)\n", endLabel)
		code += incrementSP()
	}

	_, err := c.writer.WriteString(code)

	return err
}

func pushToStack(value string) string {
	return "@SP\nA=M\nM=" + value + "\n"
}

func binaryCommandOperator(command string) (string, error) {
	switch command {
	case "add":
		return "+", nil
	case "and":
		return "&", nil
	case "or":
		return "|", nil
	default:
		return "", fmt.Errorf("%s is not a valid binary command", command)
	}
}

func unaryCommandOperator(command string) (string, error) {
	switch command {
	case "neg":
		return "-", nil
	case "not":
		return "!", nil
	default:
		return "", fmt.Errorf("%s is not a valid unary command", command)
	}
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

	code := popFromStack() + "D=M\n"
	switch segment {
	case "constant":
		return ""
	case "pointer":
		pointerAddr := "THIS"
		if index == 1 {
			pointerAddr = "THAT"
		}

		code += fmt.Sprintf("@%s\nM=D\n", pointerAddr)

		return code
	case "temp":
		addr := 5 + index
		code += fmt.Sprintf("@R%d\nM=D\n", addr)

		return code
	case "static":
		code += fmt.Sprintf("@%s.%d\nM=D\n", c.namespace, index)

		return code
	default:
		code += fmt.Sprintf("@%s\nA=M\n", segmentAddr) + incrementAddr(index) + "M=D\n"

		return code
	}
}

func popFromStack() string {
	return "@SP\nM=M-1\nA=M\n"
}

func incrementAddr(index int) string {
	code := ""
	for i := 0; i < index; i++ {
		code += "A=A+1\n"
	}

	return code
}

func incrementSP() string {
	return "@SP\nM=M+1\n"
}

func (c *CodeWriter) Save() {
	f, err := os.Create(c.filename)
	if err != nil {
		panic(err)
	}

	f.Write(c.writer.Bytes())
}
