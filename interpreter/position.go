package interpreter

type Position struct {
	Index  int
	Line   int
	Column int
}

func NewPosition() Position {
	return Position{0, 1, 0}
}

func (p *Position) Advance(currentChar string) {
	p.Index++
	p.Column++
	if currentChar == "\n" {
		p.Line++
		p.Column = 0
	}
}
