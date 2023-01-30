package code

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
