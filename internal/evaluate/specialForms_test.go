package evaluate

import (
	"testing"

	"github.com/yashbaddi/golisp/internal/node"
)

func TestDefineSpecialForm(t *testing.T) {
	t.Run("Define integer variable", func(t *testing.T) {
		env := NewGlobalEnv()
		expr := node.List{node.Symbol("define"), node.Symbol("x"), 42}
		res, err := Eval(expr, env)
		if err != nil {
			t.Fatalf("unexpected error defining x: %v", err)
		}
		if res != 42 {
			t.Errorf("expected define to return 42, got %v", res)
		}

		// Evaluate symbol x
		resX, err := Eval(node.Symbol("x"), env)
		if err != nil {
			t.Fatalf("unexpected error evaluating x: %v", err)
		}
		if resX != 42 {
			t.Errorf("expected x to be 42, got %v", resX)
		}
	})

	t.Run("Define with expression", func(t *testing.T) {
		env := NewGlobalEnv()
		expr := node.List{
			node.Symbol("define"),
			node.Symbol("y"),
			node.List{node.Symbol("+"), 10, 20},
		}
		res, err := Eval(expr, env)
		if err != nil {
			t.Fatalf("unexpected error defining y: %v", err)
		}
		if res != 30 {
			t.Errorf("expected define to return 30, got %v", res)
		}

		// Use defined variable in arithmetic
		useExpr := node.List{node.Symbol("+"), node.Symbol("y"), 5}
		resUse, err := Eval(useExpr, env)
		if err != nil {
			t.Fatalf("unexpected error evaluating (+ y 5): %v", err)
		}
		if resUse != 35 {
			t.Errorf("expected 35, got %v", resUse)
		}
	})

	t.Run("Define error cases", func(t *testing.T) {
		env := NewGlobalEnv()

		// Non-symbol target
		_, err := Eval(node.List{node.Symbol("define"), 123, 456}, env)
		if err == nil {
			t.Errorf("expected error when defining non-symbol target, got nil")
		}

		// Incorrect argument count (1 arg)
		_, err = Eval(node.List{node.Symbol("define"), node.Symbol("a")}, env)
		if err == nil {
			t.Errorf("expected error for missing define value, got nil")
		}

		// Incorrect argument count (3 args)
		_, err = Eval(node.List{node.Symbol("define"), node.Symbol("a"), 1, 2}, env)
		if err == nil {
			t.Errorf("expected error for too many define args, got nil")
		}
	})
}

func TestIfSpecialForm(t *testing.T) {
	env := NewGlobalEnv()

	t.Run("Truthy condition evaluates then-branch", func(t *testing.T) {
		expr := node.List{
			node.Symbol("if"),
			node.List{node.Symbol(">"), 5, 2},
			100,
			200,
		}
		res, err := Eval(expr, env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != 100 {
			t.Errorf("expected 100, got %v", res)
		}
	})

	t.Run("Falsy condition evaluates else-branch", func(t *testing.T) {
		expr := node.List{
			node.Symbol("if"),
			node.List{node.Symbol("<"), 5, 2},
			100,
			200,
		}
		res, err := Eval(expr, env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != 200 {
			t.Errorf("expected 200, got %v", res)
		}
	})

	t.Run("Falsy condition without else-branch returns nil", func(t *testing.T) {
		expr := node.List{
			node.Symbol("if"),
			node.List{node.Symbol("="), 1, 2},
			100,
		}
		res, err := Eval(expr, env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != nil {
			t.Errorf("expected nil, got %v", res)
		}
	})

	t.Run("Short circuiting - unchosen branch is not evaluated", func(t *testing.T) {
		// (if true 10 (/ 1 0)) - division by zero branch should NOT be evaluated
		expr := node.List{
			node.Symbol("if"),
			true,
			10,
			node.List{node.Symbol("/"), 1, 0},
		}
		res, err := Eval(expr, env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if res != 10 {
			t.Errorf("expected 10, got %v", res)
		}
	})

	t.Run("If error cases", func(t *testing.T) {
		// Too few arguments (1 arg)
		_, err := Eval(node.List{node.Symbol("if"), true}, env)
		if err == nil {
			t.Errorf("expected error for if with 1 arg, got nil")
		}

		// Too many arguments (4 args)
		_, err = Eval(node.List{node.Symbol("if"), true, 1, 2, 3}, env)
		if err == nil {
			t.Errorf("expected error for if with 4 args, got nil")
		}
	})
}
