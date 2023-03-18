package interpreter

import "fmt"

type TokenType string

const DIGITS string = "0123456789"

var TOKENS map[string]TokenType = map[string]TokenType{"+": T_PLUS, "-": T_MINUS, "*": T_MUL, "/": T_DIV, "(": T_LPAREN, ")": T_RPAREN}

const (
	T_INT    TokenType = "T_INT"
	T_FLOAT  TokenType = "T_FLOAT"
	T_PLUS   TokenType = "T_PLUS"
	T_MINUS  TokenType = "T_MINUS"
	T_MUL    TokenType = "T_MUL"
	T_DIV    TokenType = "T_DIV"
	T_LPAREN TokenType = "T_LPAREN"
	T_RPAREN TokenType = "T_RPAREN"
)

type Token struct {
	Value string
	Type  TokenType
}

func (t Token) String() string {
	return fmt.Sprintf("Token(%s, '%s' )", t.Type, t.Value)
}
