package parser

import (
	"fmt"
)

type ParserError struct {
	tokenLiteral string
	Msg          string
}

func (e ParserError) Error() string {
	return fmt.Sprintf("Parser Error: %v: %s", e.tokenLiteral, e.Msg)
}
