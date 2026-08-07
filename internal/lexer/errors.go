package lexer

import "fmt"

type LexerError struct {
	Pos int
	Msg string
}

func (e LexerError) Error() string {
	return fmt.Sprintf("Lexical Error: %d: %s", e.Pos, e.Msg)
}
