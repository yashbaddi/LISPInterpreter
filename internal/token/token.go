package token

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
