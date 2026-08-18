package parser

import (
	"strconv"

	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/node"
	"github.com/yashbaddi/golisp/internal/token"
)

type Parser struct {
	lexer *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		lexer: l,
	}
}

func (p *Parser) Parse() (any, error) {
	err := p.newToken()
	if err != nil {
		return nil, ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Bad Token",
		}
	}
	return p.parseExpression()
}

func (p *Parser) parseExpression() (any, error) {
	switch p.currToken.Type {
	case token.LPAREN:
		return p.parseList()
	case token.NUMBER:
		return p.parseNumber()
	case token.IDENTIFIER:
		return p.parseSymbol()
	case token.STRING:
		return p.currToken.Literal, nil
	case token.RPAREN:
		return nil, ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Unexpected Close Parenthesis",
		}
	default:
		return nil, ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Invalid Token",
		}
	}
}

func (p *Parser) newToken() error {
	var err error
	p.currToken, err = p.lexer.NextToken()
	if err != nil {
		return ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Failed to fetch a new token",
		}
	}
	p.peekToken, err = p.lexer.Peek()
	if err != nil {
		return ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Peek Failed",
		}
	}
	return nil
}

func (p *Parser) parseSymbol() (node.Symbol, error) {
	return node.Symbol(p.currToken.Literal), nil
}

func (p *Parser) parseNumber() (int, error) {
	val, err := strconv.Atoi(p.currToken.Literal)
	if err != nil {
		return 0, ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Invalid Number",
		}
	}
	return val, nil
}

func (p *Parser) parseList() (node.List, error) {
	var l node.List
	for p.peekToken.Type != token.RPAREN {
		if p.peekToken.Type == token.EOF {
			return nil, ParserError{
				tokenLiteral: p.currToken.Literal,
				Msg:          "Unexpected End of Token",
			}
		}
		err := p.newToken()
		if err != nil {
			return nil, err
		}
		expr, err := p.parseExpression()
		if err != nil {
			return nil, err
		}
		l = append(l, expr)
	}
	err := p.newToken()
	if err != nil {
		return nil, err
	}
	return l, nil
}
