package evaluate

import (
	"testing"

	"github.com/yashbaddi/golisp/internal/node"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr     any
		expected any
	}{
		{123, 123},
		{node.List{node.Symbol("+"), 1, 2}, 3},
		{node.List{node.Symbol("+"), 10, node.List{node.Symbol("+"), 20, 30}}, 60},
		{node.List{node.Symbol("-"), 10, 3}, 7},
	}

	for _, tt := range tests {
		env := NewGlobalEnv()
		res, err := Eval(tt.expr, env)
		if err != nil {
			t.Fatalf("Eval(%v) returned error: %v", tt.expr, err)
		}
		if res != tt.expected {
			t.Errorf("Eval(%v) = %v; want %v", tt.expr, res, tt.expected)
		}
	}
}
