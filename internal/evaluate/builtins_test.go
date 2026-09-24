package evaluate

import (
	"math"
	"reflect"
	"testing"

	"github.com/yashbaddi/golisp/internal/node"
)

func TestBuiltinFunctions(t *testing.T) {
	env := NewGlobalEnv()

	tests := []struct {
		name     string
		expr     any
		expected any
	}{
		// Arithmetic
		{"multiplication", node.List{node.Symbol("*"), 2, 3, 4}, 24},
		{"multiplication float", node.List{node.Symbol("*"), 2.5, 2}, 5.0},
		{"division exact", node.List{node.Symbol("/"), 10, 2}, 5},
		{"division float", node.List{node.Symbol("/"), 10, 4}, 2.5},
		{"division unary", node.List{node.Symbol("/"), 2.0}, 0.5},
		{"mod", node.List{node.Symbol("mod"), 10, 3}, 1},
		{"incf default", node.List{node.Symbol("incf"), 5}, 6},
		{"incf delta", node.List{node.Symbol("incf"), 5, 3}, 8},
		{"decf default", node.List{node.Symbol("decf"), 10}, 9},
		{"decf delta", node.List{node.Symbol("decf"), 10, 4}, 6},
		{"abs int", node.List{node.Symbol("abs"), -42}, 42},
		{"abs float", node.List{node.Symbol("abs"), -3.14}, 3.14},
		{"pow", node.List{node.Symbol("pow"), 2, 3}, 8},

		// Constants
		{"pi constant", node.Symbol("pi"), math.Pi},

		// Comparisons
		{"greaterThan true", node.List{node.Symbol(">"), 5, 3, 1}, true},
		{"greaterThan false", node.List{node.Symbol(">"), 5, 5}, false},
		{"greaterOrEqual true", node.List{node.Symbol(">="), 5, 5, 2}, true},
		{"lessThan true", node.List{node.Symbol("<"), 1, 3, 5}, true},
		{"lessOrEqual true", node.List{node.Symbol("<="), 1, 1, 3}, true},
		{"equalNum true", node.List{node.Symbol("="), 4, 4, 4}, true},
		{"equalNum false", node.List{node.Symbol("="), 4, 4, 5}, false},
		{"equal? true", node.List{node.Symbol("equal?"), node.List{node.Symbol("list"), 1, 2}, node.List{node.Symbol("list"), 1, 2}}, true},
		{"equal? false", node.List{node.Symbol("equal?"), node.List{node.Symbol("list"), 1, 2}, node.List{node.Symbol("list"), 1, 3}}, false},

		// List operations
		{"list", node.List{node.Symbol("list"), 1, 2, 3}, node.List{1, 2, 3}},
		{"cons", node.List{node.Symbol("cons"), 0, node.List{node.Symbol("list"), 1, 2}}, node.List{0, 1, 2}},
		{"car", node.List{node.Symbol("car"), node.List{node.Symbol("list"), 10, 20, 30}}, 10},
		{"cdr", node.List{node.Symbol("cdr"), node.List{node.Symbol("list"), 10, 20, 30}}, node.List{20, 30}},
		{"map abs", node.List{node.Symbol("map"), node.Symbol("abs"), node.List{node.Symbol("list"), -1, -2, 3}}, node.List{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Eval(tt.expr, env)
			if err != nil {
				t.Fatalf("Eval(%v) failed with error: %v", tt.expr, err)
			}
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Eval(%v) = %v (%T); want %v (%T)", tt.expr, res, res, tt.expected, tt.expected)
			}
		})
	}
}
