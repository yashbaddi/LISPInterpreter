package ast_test

import (
	"testing"

	"github.com/yashbaddi/golisp/internal/ast"
)

func TestASTString(t *testing.T) {
	program := ast.Program{
		Expression: ast.List{
			Elements: []ast.Expression{
				ast.Identifier{Value: "+"},
				ast.NumberLiteral{Value: 1},
				ast.List{
					Elements: []ast.Expression{
						ast.Identifier{Value: "*"},
						ast.NumberLiteral{Value: 2},
						ast.NumberLiteral{Value: 3},
					},
				},
			},
		},
	}

	expected := "(+ 1 (* 2 3))"
	if program.String() != expected {
		t.Errorf("program.String() wrong. got=%q, want=%q", program.String(), expected)
	}
}
