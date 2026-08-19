package lexer

import (
	"strconv"
	"strings"
	"unicode"
)

func (l *Lexer) isStartingOfNumber() bool {
	if l.eof() {
		return false
	}
	ch := l.current()
	if unicode.IsDigit(ch) {
		return true
	}

	rest := l.input[l.pos:]
	if (ch == '+' || ch == '-') && len(rest) > 1 {
		if unicode.IsDigit(rest[1]) {
			return true
		}
		if rest[1] == '.' && len(rest) > 2 && unicode.IsDigit(rest[2]) {
			return true
		}
	}
	if ch == '.' && len(rest) > 1 && unicode.IsDigit(rest[1]) {
		return true
	}
	return false
}

func isSymbolOperator(r rune) bool {
	return strings.ContainsRune("+-*/%><=?_", r)
}

func (l *Lexer) readSymbolOperator() string {
	start := l.pos
	for !l.eof() && isSymbolOperator(l.current()) {
		l.advance()
	}
	return string(l.input[start:l.pos])
}

func (l *Lexer) readNumber() (string, error) {
	var builder strings.Builder
	start := l.pos
	hasDecimal := false

	if !l.eof() && (l.current() == '-' || l.current() == '+') {
		builder.WriteRune(l.current())
		l.advance()
	}

	for !l.eof() {
		r := l.current()

		if unicode.IsDigit(r) {
			builder.WriteRune(r)
			l.advance()
		} else if r == '.' && !hasDecimal {
			hasDecimal = true
			builder.WriteRune(r)
			l.advance()
		} else {
			break
		}
	}

	result := builder.String()
	if builder.Len() == 0 || result == "-" || result == "+" || result == "." || result == "-." || result == "+." {
		return "", LexerError{
			Pos: start,
			Msg: "expected a number",
		}
	}

	return result, nil
}

func (l *Lexer) readInteger() (string, error) {
	return l.readNumber()
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
	for !l.eof() && (unicode.IsLetter(l.current()) || l.current() == '?' || l.current() == '!') {
		l.advance()
	}
	return string(l.input[start:l.pos]), nil
}
