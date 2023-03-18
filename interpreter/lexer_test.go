package interpreter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	cases := []struct {
		input string
		want  []Token
	}{
		{
			input: "1 + 2",
			want: []Token{
				{Value: "1", Type: T_INT},
				{Value: "+", Type: T_PLUS},
				{Value: "2", Type: T_INT},
			},
		},
		{
			input: "1 - (3 * 4) / 2",
			want: []Token{
				{Value: "1", Type: T_INT},
				{Value: "-", Type: T_MINUS},
				{Value: "(", Type: T_LPAREN},
				{Value: "3", Type: T_INT},
				{Value: "*", Type: T_MUL},
				{Value: "4", Type: T_INT},
				{Value: ")", Type: T_RPAREN},
				{Value: "/", Type: T_DIV},
				{Value: "2", Type: T_INT},
			},
		},
		{
			input: "3.14 + 2.0",
			want: []Token{
				{Value: "3.14", Type: T_FLOAT},
				{Value: "+", Type: T_PLUS},
				{Value: "2.0", Type: T_FLOAT},
			},
		},
	}

	for _, c := range cases {
		lex := NewLexer(c.input)
		got := lex.CreateTokens()
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("Lexer.CreateTokens() for '%s' returned %+v, want %+v", c.input, got, c.want)
		}

	}
	fmt.Println("Lexer tests passed")
}
