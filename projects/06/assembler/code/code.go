package code

import "errors"

type Code struct{}

func New() *Code {
	return &Code{}
}

func (c *Code) Dest(token string) byte {
	switch token {
	case "M":
		return 0b001
	case "D":
		return 0b010
	case "MD":
		return 0b011
	case "A":
		return 0b100
	case "AM":
		return 0b101
	case "AD":
		return 0b110
	case "AMD":
		return 0b111
	default:
		return 0b000
	}
}

func (c *Code) Comp(token string) (byte, error) {
	switch token {
	case "0":
		return 0b0101010, nil
	case "1":
		return 0b0111111, nil
	case "-1":
		return 0b0111010, nil
	case "D":
		return 0b0001100, nil
	case "A":
		return 0b0110000, nil
	case "!D":
		return 0b0001101, nil
	case "!A":
		return 0b0110001, nil
	case "-D":
		return 0b0001111, nil
	case "-A":
		return 0b0110011, nil
	case "D+1":
		return 0b0011111, nil
	case "A+1":
		return 0b0110111, nil
	case "D-1":
		return 0b0001110, nil
	case "A-1":
		return 0b0110010, nil
	case "D+A":
		return 0b0000010, nil
	case "D-A":
		return 0b0010011, nil
	case "A-D":
		return 0b0000111, nil
	case "D&A":
		return 0b0000000, nil
	case "D|A":
		return 0b0010101, nil
	case "M":
		return 0b1110000, nil
	case "!M":
		return 0b1110001, nil
	case "-M":
		return 0b1110011, nil
	case "M+1":
		return 0b1110111, nil
	case "M-1":
		return 0b1110010, nil
	case "D+M":
		return 0b1000010, nil
	case "D-M":
		return 0b1010011, nil
	case "M-D":
		return 0b1000111, nil
	case "D&M":
		return 0b1000000, nil
	case "D|M":
		return 0b1010101, nil
	default:
		return 0b0, errors.New("invalid comp detected")
	}
}

func (c *Code) Jump(token string) byte {
	switch token {
	case "JGT":
		return 0b001
	case "JEQ":
		return 0b010
	case "JGE":
		return 0b011
	case "JLT":
		return 0b100
	case "JNE":
		return 0b101
	case "JLE":
		return 0b110
	case "JMP":
		return 0b111
	default:
		return 0b0
	}
}
