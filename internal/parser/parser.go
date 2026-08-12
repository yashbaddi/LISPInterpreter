package parser

import (
	"strconv"

	"github.com/yashbaddi/golisp/internal/ast"
	"github.com/yashbaddi/golisp/internal/lexer"
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

func (p *Parser) Parse() (*ast.Program, error) {

	err := p.newToken()

	if err != nil {
		return &ast.Program{}, ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Bad Token",
		}
	}

	switch p.currToken.Type {
	case token.LPAREN:
		plist, err := p.parseList()
		if err != nil {
			return &ast.Program{}, err
		}

		return &ast.Program{
			Expression: plist,
		}, nil
	case token.NUMBER:
		pnum, err := p.parseNumber()
		if err != nil {
			return &ast.Program{}, err
		}

		return &ast.Program{
			Expression: pnum,
		}, nil
	case token.IDENTIFIER:
		pident, err := p.parseIdentifer()
		if err != nil {
			return &ast.Program{}, err
		}

		return &ast.Program{
			Expression: pident,
		}, nil
	default:
		return &ast.Program{}, ParserError{
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

func (p *Parser) parseIdentifer() (ast.Identifier, error) {
	return ast.Identifier{
		Value: p.currToken.Literal,
	}, nil
}

func (p *Parser) parseNumber() (ast.NumberLiteral, error) {
	val, err := strconv.Atoi(p.currToken.Literal)
	if err != nil {
		return ast.NumberLiteral{}, ParserError{
			tokenLiteral: p.currToken.Literal,
			Msg:          "Invalid Number",
		}
	}
	return ast.NumberLiteral{
		Value: val,
	}, nil
}

func (p *Parser) parseList() (*ast.List, error) {

	var l []ast.Expression
	for p.peekToken.Type != token.RPAREN {
		expr, err := p.Parse()
		if err != nil {
			return &ast.List{}, ParserError{
				tokenLiteral: p.currToken.Literal,
				Msg:          "Parse Fail",
			}
		}
		l = append(l, expr.Expression)
	}
	return &ast.List{
		Elements: l,
	}, nil
}
