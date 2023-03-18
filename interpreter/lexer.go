package interpreter

import (
	"fmt"
	"strings"
)

type Lexer struct {
	Text         string
	Pos          int
	CurrentChar  string
	ReservedKeys map[string]TokenType
}

func NewLexer(text string) *Lexer {
	l := &Lexer{
		Text:         text,
		ReservedKeys: TOKENS,
	}
	l.CurrentChar = string(l.Text[l.Pos])
	return l
}

func (l *Lexer) Advance() {
	l.Pos++
	if l.Pos >= len(l.Text) {
		l.CurrentChar = ""
	} else {
		l.CurrentChar = string(l.Text[l.Pos])
	}
}

func (l *Lexer) SkipWhitespace() {
	for l.CurrentChar != "" && l.CurrentChar == " " {
		l.Advance()
	}
}

func (l *Lexer) NextToken() Token {
	for l.CurrentChar != "" {
		if l.CurrentChar == " " {
			l.SkipWhitespace()
			continue
		}
		if strings.ContainsAny(l.CurrentChar, DIGITS) {
			return l.numeric()
		}

		if t, ok := l.ReservedKeys[l.CurrentChar]; ok {
			tokenValue := l.CurrentChar
			l.Advance()
			return Token{tokenValue, t}
		}
		panic(fmt.Sprintf("Invalid character: %s", l.CurrentChar))
	}
	return Token{"", ""}
}

func (l *Lexer) numeric() Token {
	result := ""
	for l.CurrentChar != "" && strings.ContainsAny(l.CurrentChar, DIGITS) {
		result += l.CurrentChar
		l.Advance()
	}
	if l.CurrentChar == "." {
		result += l.CurrentChar
		l.Advance()
		for l.CurrentChar != "" && strings.ContainsAny(l.CurrentChar, DIGITS) {
			result += l.CurrentChar
			l.Advance()
		}
		return Token{result, "T_FLOAT"}
	}
	return Token{result, "T_INT"}
}

func (l *Lexer) CreateTokens() []Token {
	var tokens []Token
	for token := l.NextToken(); token.Type != ""; token = l.NextToken() {
		tokens = append(tokens, token)
	}
	return tokens
}
