package token

import (
	"fmt"
	"slices"
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
	input string
	pos   int
}

type LexerError struct {
	Pos int
	Msg string
}

func (e LexerError) Error() string {
	return fmt.Sprintf("Lexical Error: %d: %s", e.Pos, e.Msg)
}

func (l *Lexer) NextToken() (Token, error) {

	l.skipWhiteSpace()
	if l.pos >= len(l.input) {
		return Token{Type: EOF, Literal: ""}, nil
	}

	ch := rune(l.input[l.pos])
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
		return Token{Type: STRING, Literal: str}, err
	case unicode.IsDigit(ch):
		num, err := l.readNumber()
		return Token{Type: NUMBER, Literal: num}, err
	case unicode.IsLetter(ch):
		idet, err := l.readIdentifier()
		return Token{Type: IDENTIFIER, Literal: idet}, err
	default:
		return Token{Type: ILLEGAL, Literal: string(ch)}, LexerError{Pos: l.pos, Msg: "Invalid Token Encountered"}
	}

}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(rune(l.input[l.pos])) {
		l.pos += 1
	}
}

func (l *Lexer) readNumber() (string, error) {
	var builder strings.Builder

	if l.pos >= len(l.input) {
		return "", LexerError{
			Pos: l.pos,
			Msg: "unexpected end of input",
		}
	}

	for l.pos < len(l.input) {
		r := rune(l.input[l.pos])

		if !unicode.IsDigit(r) {
			break
		}

		builder.WriteRune(r)
		l.pos++
	}

	if builder.Len() == 0 {
		return "", LexerError{
			Pos: l.pos,
			Msg: "expected a number",
		}
	}

	return builder.String(), nil
}

func (l *Lexer) readString() (string, error) {
	var builder strings.Builder

	for l.pos < len(l.input) {
		switch l.input[l.pos] {
		case '"':
			l.pos++
			return builder.String(), nil
		case '\\':
			l.pos++
			if l.pos == len(l.input) {
				return "", LexerError{Pos: l.pos, Msg: "Unexpected End of file"}
			}
			selectedRunes := []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'}
			r := rune(l.input[l.pos])
			if slices.Contains(selectedRunes, r) {
				l.pos++
				builder.WriteRune(r)
				continue
			}
			if r == 'u' {
				if len(l.input)-l.pos < 5 {
					return "", LexerError{Pos: l.pos, Msg: "Invalid Hex Value"}
				}

				hex := l.input[l.pos+1 : l.pos+5]

				value, err := strconv.ParseUint(hex, 16, 16)
				if err != nil {
					return "", LexerError{Pos: l.pos, Msg: "Invalid Hex Value"}
				}

				l.pos += 5
				builder.WriteRune(rune(value))
				continue
			}
		default:
			builder.WriteByte(l.input[l.pos])
			l.pos++
		}
	}
	return "", LexerError{Pos: l.pos, Msg: "Unexpected End of file"}

}

func (l *Lexer) readIdentifier() (string, error) {
	start := l.pos
	for unicode.IsLetter(rune(l.input[l.pos])) {
		l.pos++
	}
	return l.input[start:l.pos], nil
}
