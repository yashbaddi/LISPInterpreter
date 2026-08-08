package lexer

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/yashbaddi/golisp/internal/token"
)

type Lexer struct {
	input []rune
	pos   int
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: []rune(input),
	}
}

func (l *Lexer) NextToken() (token.Token, error) {

	l.skipWhiteSpace()
	if l.eof() {
		return token.Token{Type: token.EOF, Literal: ""}, nil
	}

	ch := l.current()
	switch {
	case ch == '(':
		l.advance()
		return token.Token{Type: token.LPAREN, Literal: "("}, nil
	case ch == ')':
		l.advance()
		return token.Token{Type: token.RPAREN, Literal: ")"}, nil
	case ch == '"':
		l.advance()
		str, err := l.readString()
		if err != nil {
			return token.Token{}, err
		}
		return token.Token{Type: token.STRING, Literal: str}, nil
	case unicode.IsDigit(ch):
		num, err := l.readInteger()
		if err != nil {
			return token.Token{}, err
		}
		return token.Token{Type: token.NUMBER, Literal: num}, nil
	case unicode.IsLetter(ch):
		idet, err := l.readIdentifier()
		if err != nil {
			return token.Token{}, err
		}
		return token.Token{Type: token.IDENTIFIER, Literal: idet}, nil
	default:
		return token.Token{}, LexerError{Pos: l.pos, Msg: "Invalid Token Encountered"}
	}

}

func (l *Lexer) skipWhiteSpace() {
	for !l.eof() && unicode.IsSpace(l.input[l.pos]) {
		l.advance()
	}
}

func (l *Lexer) readInteger() (string, error) {
	var builder strings.Builder
	start := l.pos
	for !l.eof() {
		r := l.current()

		if !unicode.IsDigit(r) {
			break
		}

		builder.WriteRune(r)
		l.advance()
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

	for !l.eof() {
		switch l.current() {
		case '"':
			l.advance()
			return builder.String(), nil
		case '\\':
			l.advance()
			if l.eof() {
				return "", LexerError{Pos: start, Msg: "Unexpected End of file"}
			}

			r := l.current()
			l.advance()

			switch r {
			case '"':
				builder.WriteRune('"')
			case '\\':
				builder.WriteRune('\\')
			case '/':
				builder.WriteRune('/')
			case 'b':
				builder.WriteRune('\b')
			case 'f':
				builder.WriteRune('\f')
			case 'n':
				builder.WriteRune('\n')
			case 'r':
				builder.WriteRune('\r')
			case 't':
				builder.WriteRune('\t')
			case 'u':
				if len(l.input)-l.pos < 4 {
					return "", LexerError{Pos: start, Msg: "Invalid Hex Value"}
				}

				hex := string(l.input[l.pos : l.pos+4])

				value, err := strconv.ParseUint(hex, 16, 16)
				if err != nil {
					return "", LexerError{Pos: start, Msg: "Invalid Hex Value"}
				}

				l.advanceN(4)

				builder.WriteRune(rune(value))
			default:
				return "", LexerError{
					Pos: start,
					Msg: "invalid escape sequence",
				}
			}
		default:
			builder.WriteRune(l.current())
			l.advance()
		}
	}
	return "", LexerError{Pos: start, Msg: "Unexpected End of file"}

}

func (l *Lexer) readIdentifier() (string, error) {
	start := l.pos
	for !l.eof() && unicode.IsLetter(l.current()) {
		l.advance()
	}
	return string(l.input[start:l.pos]), nil
}

func (l *Lexer) current() rune {
	return l.input[l.pos]
}

func (l *Lexer) advance() {
	l.pos++
}

func (l *Lexer) advanceN(n int) {
	l.pos += n
	if l.pos > len(l.input) {
		l.pos = len(l.input)
	}
}

func (l *Lexer) eof() bool {
	return l.pos >= len(l.input)
}
