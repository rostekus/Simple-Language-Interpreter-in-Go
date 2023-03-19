package interpreter

import "fmt"

type NumericNode struct {
	Token Token
}

func (n *NumericNode) String() string {
	return n.Token.String()
}

type BinaryOperatorNode struct {
	LeftNode  Token
	Operator  Token
	RightNode Token
}

func (n *BinaryOperatorNode) String() string {
	return fmt.Sprintf("%s %s %s", n.LeftNode, n.Operator, n.RightNode)
}
