package token

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type TokenType int

const (
	LPAREN TokenType = iota
	RPAREN

	STRING
	NUMBER
	IDENTIFIER

	ILLEGAL
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}

type Lexer struct {
	input []rune
	pos   int
}

type LexerError struct {
	Pos int
	Msg string
}

func (e LexerError) Error() string {
	return fmt.Sprintf("Lexical Error: %d: %s", e.Pos, e.Msg)
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
	}
}

func (l *Lexer) NextToken() (Token, error) {

	l.skipWhiteSpace()
	if l.pos >= len(l.input) {
		return Token{Type: EOF, Literal: ""}, nil
	}

	ch := l.input[l.pos]
	switch {
	case ch == '(':
		l.pos++
		return Token{Type: LPAREN, Literal: "("}, nil
	case ch == ')':
		l.pos++
		return Token{Type: RPAREN, Literal: ")"}, nil
	case ch == '"':
		l.pos++
		str, err := l.readString()
		if err != nil {
			return Token{}, err
		}
		return Token{Type: STRING, Literal: str}, nil
	case unicode.IsDigit(ch):
		num, err := l.readInteger()
		if err != nil {
			return Token{}, err
		}
		return Token{Type: NUMBER, Literal: num}, nil
	case unicode.IsLetter(ch):
		idet, err := l.readIdentifier()
		if err != nil {
			return Token{}, err
		}
		return Token{Type: IDENTIFIER, Literal: idet}, nil
	default:
		return Token{}, LexerError{Pos: l.pos, Msg: "Invalid Token Encountered"}
	}

}

func (l *Lexer) skipWhiteSpace() {
	for l.pos < len(l.input) && unicode.IsSpace(l.input[l.pos]) {
		l.pos += 1
	}
}

func (l *Lexer) readInteger() (string, error) {
	var builder strings.Builder
	start := l.pos
	for l.pos < len(l.input) {
		r := l.input[l.pos]

		if !unicode.IsDigit(r) {
			break
		}

		builder.WriteRune(r)
		l.pos++
	}

	if builder.Len() == 0 {
		return "", LexerError{
			Pos: start,
			Msg: "expected a number",
		}
	}

	return builder.String(), nil
}

func (l *Lexer) readString() (string, error) {
	var builder strings.Builder
	start := l.pos

	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case '"':
			l.pos++
			return builder.String(), nil
		case '\\':
			l.pos++
			if l.pos == len(l.input) {
				return "", LexerError{Pos: start, Msg: "Unexpected End of file"}
			}

			r := l.input[l.pos]

			switch r {
			case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
				l.pos++
				builder.WriteRune(r)
				continue

			case 'u':
				if len(l.input)-l.pos < 5 {
					return "", LexerError{Pos: start, Msg: "Invalid Hex Value"}
				}

				hex := string(l.input[l.pos+1 : l.pos+5])

				value, err := strconv.ParseUint(hex, 16, 16)
				if err != nil {
					return "", LexerError{Pos: start, Msg: "Invalid Hex Value"}
				}

				l.pos += 5
				builder.WriteRune(rune(value))
				continue
			}
		default:
			builder.WriteRune(l.input[l.pos])
			l.pos++
		}
	}
	return "", LexerError{Pos: start, Msg: "Unexpected End of file"}

}

func (l *Lexer) readIdentifier() (string, error) {
	start := l.pos
	for l.pos < len(l.input) && unicode.IsLetter(l.input[l.pos]) {
		l.pos++
	}
	return string(l.input[start:l.pos]), nil
}
