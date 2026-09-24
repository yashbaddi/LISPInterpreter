package tests

import (
	"reflect"
	"testing"

	"github.com/yashbaddi/golisp/internal/evaluate"
	"github.com/yashbaddi/golisp/internal/lexer"
	"github.com/yashbaddi/golisp/internal/node"
	"github.com/yashbaddi/golisp/internal/parser"
)

func runLisp(input string, env *evaluate.Env) (any, error) {
	l := lexer.NewLexer(input)
	p := parser.NewParser(l)
	expr, err := p.Parse()
	if err != nil {
		return nil, err
	}
	return evaluate.Eval(expr, env)
}

func TestBasicArithmetic(t *testing.T) {
	env := evaluate.NewGlobalEnv()

	t.Run("should add numbers correctly", func(t *testing.T) {
		res1, err := runLisp("(+ 1 1)", env)
		if err != nil || res1 != 2 {
			t.Errorf("expected 2, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("(+ 1 2 3 4)", env)
		if err != nil || res2 != 10 {
			t.Errorf("expected 10, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should subtract numbers correctly", func(t *testing.T) {
		res1, err := runLisp("(- 10 3)", env)
		if err != nil || res1 != 7 {
			t.Errorf("expected 7, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("(- 10 2 1)", env)
		if err != nil || res2 != 7 {
			t.Errorf("expected 7, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should multiply numbers correctly", func(t *testing.T) {
		res, err := runLisp("(* 2 3 4)", env)
		if err != nil || res != 24 {
			t.Errorf("expected 24, got %v (err: %v)", res, err)
		}
	})

	t.Run("should handle division", func(t *testing.T) {
		res, err := runLisp("(/ 10 2)", env)
		if err != nil || res != 5 {
			t.Errorf("expected 5, got %v (err: %v)", res, err)
		}

		res2, err := runLisp("(/ 100 10 2)", env)
		if err != nil || res2 != 5 {
			t.Errorf("expected 5, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should handle mod", func(t *testing.T) {
		res1, err := runLisp("(mod 10 3)", env)
		if err != nil || res1 != 1 {
			t.Errorf("expected 1, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("(mod 3 10)", env)
		if err != nil || res2 != 3 {
			t.Errorf("expected 3, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should handle power/exponentiation", func(t *testing.T) {
		res1, err := runLisp("(pow 2 10)", env)
		if err != nil || res1 != 1024 {
			t.Errorf("expected 1024, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("(pow 2 16)", env)
		if err != nil || res2 != 65536 {
			t.Errorf("expected 65536, got %v (err: %v)", res2, err)
		}
	})
}

func TestVariableDefinitionsAndScope(t *testing.T) {
	env := evaluate.NewGlobalEnv()

	t.Run("should define variables", func(t *testing.T) {
		_, err := runLisp("(define x 10)", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		res, err := runLisp("(+ x 5)", env)
		if err != nil || res != 15 {
			t.Errorf("expected 15, got %v (err: %v)", res, err)
		}
	})

	t.Run("should handle nested definitions and lambdas", func(t *testing.T) {
		_, err := runLisp("(define twice (lambda (x) (* 2 x)))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		res, err := runLisp("(twice 5)", env)
		if err != nil || res != 10 {
			t.Errorf("expected 10, got %v (err: %v)", res, err)
		}
	})

	t.Run("should handle higher-order functions", func(t *testing.T) {
		_, err := runLisp("(define repeat (lambda (f) (lambda (x) (f (f x)))))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = runLisp("(define twice (lambda (x) (* 2 x)))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		res1, err := runLisp("((repeat twice) 10)", env)
		if err != nil || res1 != 40 {
			t.Errorf("expected 40, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("((repeat (repeat twice)) 10)", env)
		if err != nil || res2 != 160 {
			t.Errorf("expected 160, got %v (err: %v)", res2, err)
		}

		res3, err := runLisp("((repeat (repeat (repeat twice))) 10)", env)
		if err != nil || res3 != 2560 {
			t.Errorf("expected 2560, got %v (err: %v)", res3, err)
		}

		res4, err := runLisp("((repeat (repeat (repeat (repeat twice)))) 10)", env)
		if err != nil || res4 != 655360 {
			t.Errorf("expected 655360, got %v (err: %v)", res4, err)
		}
	})

	t.Run("should handle lexical scoping", func(t *testing.T) {
		_, err := runLisp("(define x 10)", env)
		if err != nil {
			t.Fatalf("step 1 error: %v", err)
		}
		_, err = runLisp("(define makeadder (lambda (x) (lambda (y) (+ x y))))", env)
		if err != nil {
			t.Fatalf("step 2 error: %v", err)
		}
		_, err = runLisp("(define add5 (makeadder 5))", env)
		if err != nil {
			t.Fatalf("step 3 error: %v", err)
		}

		res, err := runLisp("(add5 10)", env)
		if err != nil || res != 15 {
			t.Errorf("expected 15, got %v (err: %v)", res, err)
		}

		origX, err := runLisp("x", env)
		if err != nil || origX != 10 {
			t.Errorf("expected x to remain 10, got %v (err: %v)", origX, err)
		}
	})
}

func TestConditionalsAndComparisons(t *testing.T) {
	env := evaluate.NewGlobalEnv()

	t.Run("should handle booleans", func(t *testing.T) {
		res1, err := runLisp("#t", env)
		if err != nil || res1 != true {
			t.Errorf("expected true, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("#f", env)
		if err != nil || res2 != false {
			t.Errorf("expected false, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should handle if conditions", func(t *testing.T) {
		res1, err := runLisp("(if (> 10 5) 1 0)", env)
		if err != nil || res1 != 1 {
			t.Errorf("expected 1, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("(if (< 10 5) 1 0)", env)
		if err != nil || res2 != 0 {
			t.Errorf("expected 0, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should handle nested if conditions", func(t *testing.T) {
		res, err := runLisp("(if #t (if #f 3 2) 7)", env)
		if err != nil || res != 2 {
			t.Errorf("expected 2, got %v (err: %v)", res, err)
		}
	})

	t.Run("should handle comparison operators", func(t *testing.T) {
		tests := []struct {
			expr     string
			expected bool
		}{
			{"(>= 10 10)", true},
			{"(<= 5 10)", true},
			{"(= 42 42)", true},
			{"(equal? 100 100)", true},
		}

		for _, tt := range tests {
			res, err := runLisp(tt.expr, env)
			if err != nil || res != tt.expected {
				t.Errorf("%s = %v (err: %v); want %v", tt.expr, res, err, tt.expected)
			}
		}
	})

	t.Run("should handle recursive factorial", func(t *testing.T) {
		_, err := runLisp("(define fact (lambda (n) (if (<= n 1) 1 (* n (fact (- n 1))))))", env)
		if err != nil {
			t.Fatalf("unexpected error defining fact: %v", err)
		}

		res1, err := runLisp("(fact 5)", env)
		if err != nil || res1 != 120 {
			t.Errorf("expected 120, got %v (err: %v)", res1, err)
		}

		res2, err := runLisp("(fact 10)", env)
		if err != nil || res2 != 3628800 {
			t.Errorf("expected 3628800, got %v (err: %v)", res2, err)
		}
	})

	t.Run("should handle recursive fibonacci", func(t *testing.T) {
		_, err := runLisp("(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))", env)
		if err != nil {
			t.Fatalf("unexpected error defining fib: %v", err)
		}

		res, err := runLisp("(fib 10)", env)
		if err != nil || res != 89 {
			t.Errorf("expected 89, got %v (err: %v)", res, err)
		}
	})
}

func TestListOperations(t *testing.T) {
	env := evaluate.NewGlobalEnv()

	t.Run("should handle list creation and manipulation", func(t *testing.T) {
		_, err := runLisp("(define mylist (list 1 2 3 0 0))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		resCar, err := runLisp("(car mylist)", env)
		if err != nil || resCar != 1 {
			t.Errorf("expected 1, got %v (err: %v)", resCar, err)
		}

		resCdr, err := runLisp("(cdr mylist)", env)
		expectedCdr := node.List{2, 3, 0, 0}
		if err != nil || !reflect.DeepEqual(resCdr, expectedCdr) {
			t.Errorf("expected %v, got %v (err: %v)", expectedCdr, resCdr, err)
		}
	})

	t.Run("should handle list processing", func(t *testing.T) {
		_, err := runLisp("(define first car)", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = runLisp("(define rest cdr)", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = runLisp("(define count (lambda (item L) (if L (+ (if (equal? item (first L)) 1 0) (count item (rest L))) 0)))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		res, err := runLisp("(count 0 (list 0 1 2 3 0 0))", env)
		if err != nil || res != 3 {
			t.Errorf("expected 3, got %v (err: %v)", res, err)
		}
	})

	t.Run("should handle map and range", func(t *testing.T) {
		_, err := runLisp("(define fib (lambda (n) (if (< n 2) 1 (+ (fib (- n 1)) (fib (- n 2))))))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		_, err = runLisp("(define range (lambda (a b) (if (= a b) (list) (cons a (range (+ a 1) b)))))", env)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		res, err := runLisp("(map fib (range 0 10))", env)
		expectedFibs := node.List{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
		if err != nil || !reflect.DeepEqual(res, expectedFibs) {
			t.Errorf("expected %v, got %v (err: %v)", expectedFibs, res, err)
		}
	})
}

func TestStrings(t *testing.T) {
	env := evaluate.NewGlobalEnv()

	t.Run("should parse strings", func(t *testing.T) {
		res, err := runLisp(`"Hello World"`, env)
		if err != nil || res != "Hello World" {
			t.Errorf(`expected "Hello World", got %v (err: %v)`, res, err)
		}
	})
}
