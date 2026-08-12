package lexer

import (
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
	case ch == '+':
		l.advance()
		return token.Token{Type: token.IDENTIFIER, Literal: "+"}, nil
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

func (l *Lexer) current() rune {
	return l.input[l.pos]
}

func (l *Lexer) Peek() (token.Token, error) {
	start := l.pos

	tok, err := l.NextToken()
	l.pos = start

	if err != nil {
		return token.Token{}, LexerError{
			Pos: l.pos,
			Msg: "Peek Failed",
		}
	}
	return tok, nil
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
