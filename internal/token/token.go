package token

type TokenType int

const (
	LPAREN TokenType = iota
	RPAREN

	STRING
	NUMBER
	BOOLEAN
	IDENTIFIER

	ILLEGAL
	EOF
)

type Token struct {
	Type    TokenType
	Literal string
}
