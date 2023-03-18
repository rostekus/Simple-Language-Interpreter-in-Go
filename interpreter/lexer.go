package interpreter

import (
	"fmt"
	"strings"
)

type Lexer struct {
	Text         string
	Pos          Position
	CurrentChar  string
	ReservedKeys map[string]TokenType
}

func NewLexer(text string) *Lexer {
	l := &Lexer{
		Text:         text,
		ReservedKeys: TOKENS,
		Pos:          NewPosition(),
		CurrentChar:  string(text[0]),
	}
	return l
}

func (l *Lexer) Advance() {
	l.Pos.Advance(l.CurrentChar)
	if l.Pos.Index >= len(l.Text) {
		l.CurrentChar = ""
	} else {
		l.CurrentChar = string(l.Text[l.Pos.Index])
	}
}

func (l *Lexer) SkipWhitespace() {
	for l.CurrentChar != "" && l.CurrentChar == " " {
		l.Advance()
	}
}

func (l *Lexer) NextToken() (Token, error) {
	for l.CurrentChar != "" {
		if l.CurrentChar == " " || l.CurrentChar == "\t" || l.CurrentChar == "\n" {
			l.SkipWhitespace()
			continue
		}
		if strings.ContainsAny(l.CurrentChar, DIGITS) {
			return l.numeric(), nil
		}

		if t, ok := l.ReservedKeys[l.CurrentChar]; ok {
			tokenValue := l.CurrentChar
			l.Advance()
			return Token{tokenValue, t}, nil
		}

		return Token{}, &NameError{fmt.Sprintf("Unknown character '%s'", l.CurrentChar), l.Pos.Line, "<stdin>"}

	}
	return Token{"", ""}, nil
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
	for {
		token, err := l.NextToken()
		if err != nil {
			fmt.Println(err)
		}
		if token.Type == "" {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens
}
