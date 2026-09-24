package evaluate

import (
	"fmt"
	"math"
	"reflect"

	"github.com/yashbaddi/golisp/internal/node"
)

type Env struct {
	vars  map[node.Symbol]any
	outer *Env
}

func NewEnv(outer *Env) *Env {
	return &Env{
		vars:  make(map[node.Symbol]any),
		outer: outer,
	}
}

func (e *Env) Get(sym node.Symbol) (any, bool) {
	val, ok := e.vars[sym]
	if ok {
		return val, true
	}
	if e.outer != nil {
		return e.outer.Get(sym)
	}
	return nil, false
}

func (e *Env) Set(sym node.Symbol, val any) {
	e.vars[sym] = val
}

func toFloat(val any) (float64, error) {
	switch v := val.(type) {
	case int:
		return float64(v), nil
	case float64:
		return v, nil
	default:
		return 0, fmt.Errorf("expected number, got %T (%v)", val, val)
	}
}

func add(args []any) (any, error) {
	hasFloat := false
	floatSum := 0.0
	intSum := 0

	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			intSum += v
			floatSum += float64(v)
		case float64:
			hasFloat = true
			floatSum += v
		default:
			return nil, fmt.Errorf("invalid argument type for +: expected number, got %T", arg)
		}
	}

	if hasFloat {
		return floatSum, nil
	}
	return intSum, nil
}

func subtract(args []any) (any, error) {
	if len(args) == 0 {
		return 0, fmt.Errorf("subtraction requires at least one argument")
	}

	firstF, err := toFloat(args[0])
	if err != nil {
		return nil, err
	}

	if len(args) == 1 {
		if _, ok := args[0].(int); ok {
			return -args[0].(int), nil
		}
		return -firstF, nil
	}

	hasFloat := false
	if _, ok := args[0].(float64); ok {
		hasFloat = true
	}

	resFloat := firstF
	resInt := 0
	if firstInt, ok := args[0].(int); ok {
		resInt = firstInt
	}

	for _, arg := range args[1:] {
		switch v := arg.(type) {
		case int:
			resInt -= v
			resFloat -= float64(v)
		case float64:
			hasFloat = true
			resFloat -= v
		default:
			return nil, fmt.Errorf("invalid argument type for -: expected number, got %T", arg)
		}
	}

	if hasFloat {
		return resFloat, nil
	}
	return resInt, nil
}

func multiply(args []any) (any, error) {
	hasFloat := false
	floatProd := 1.0
	intProd := 1

	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			intProd *= v
			floatProd *= float64(v)
		case float64:
			hasFloat = true
			floatProd *= v
		default:
			return nil, fmt.Errorf("invalid argument type for *: expected number, got %T", arg)
		}
	}

	if hasFloat {
		return floatProd, nil
	}
	return intProd, nil
}

func divide(args []any) (any, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("division requires at least one argument")
	}

	firstF, err := toFloat(args[0])
	if err != nil {
		return nil, err
	}

	if len(args) == 1 {
		if firstF == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return 1.0 / firstF, nil
	}

	resFloat := firstF
	for _, arg := range args[1:] {
		f, err := toFloat(arg)
		if err != nil {
			return nil, err
		}
		if f == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		resFloat /= f
	}

	// If result is integer value and original args were ints, return int
	allInts := true
	for _, arg := range args {
		if _, ok := arg.(int); !ok {
			allInts = false
			break
		}
	}

	if allInts && resFloat == math.Floor(resFloat) {
		return int(resFloat), nil
	}
	return resFloat, nil
}

func mod(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("mod requires exactly 2 arguments, got %d", len(args))
	}
	a, ok1 := args[0].(int)
	b, ok2 := args[1].(int)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("mod arguments must be integers")
	}
	if b == 0 {
		return nil, fmt.Errorf("modulo by zero")
	}
	return a % b, nil
}

func incf(args []any) (any, error) {
	if len(args) < 1 || len(args) > 2 {
		return nil, fmt.Errorf("incf requires 1 or 2 arguments, got %d", len(args))
	}
	delta := 1
	if len(args) == 2 {
		d, ok := args[1].(int)
		if !ok {
			return nil, fmt.Errorf("incf delta must be an integer, got %T", args[1])
		}
		delta = d
	}

	switch v := args[0].(type) {
	case int:
		return v + delta, nil
	case float64:
		return v + float64(delta), nil
	default:
		return nil, fmt.Errorf("incf target must be a number, got %T", args[0])
	}
}

func decf(args []any) (any, error) {
	if len(args) < 1 || len(args) > 2 {
		return nil, fmt.Errorf("decf requires 1 or 2 arguments, got %d", len(args))
	}
	delta := 1
	if len(args) == 2 {
		d, ok := args[1].(int)
		if !ok {
			return nil, fmt.Errorf("decf delta must be an integer, got %T", args[1])
		}
		delta = d
	}

	switch v := args[0].(type) {
	case int:
		return v - delta, nil
	case float64:
		return v - float64(delta), nil
	default:
		return nil, fmt.Errorf("decf target must be a number, got %T", args[0])
	}
}

func greaterThan(args []any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("> requires at least 2 arguments")
	}
	for i := 0; i < len(args)-1; i++ {
		f1, err1 := toFloat(args[i])
		f2, err2 := toFloat(args[i+1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("> arguments must be numbers")
		}
		if !(f1 > f2) {
			return false, nil
		}
	}
	return true, nil
}

func greaterOrEqual(args []any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf(">= requires at least 2 arguments")
	}
	for i := 0; i < len(args)-1; i++ {
		f1, err1 := toFloat(args[i])
		f2, err2 := toFloat(args[i+1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf(">= arguments must be numbers")
		}
		if !(f1 >= f2) {
			return false, nil
		}
	}
	return true, nil
}

func lessThan(args []any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("< requires at least 2 arguments")
	}
	for i := 0; i < len(args)-1; i++ {
		f1, err1 := toFloat(args[i])
		f2, err2 := toFloat(args[i+1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("< arguments must be numbers")
		}
		if !(f1 < f2) {
			return false, nil
		}
	}
	return true, nil
}

func lessOrEqual(args []any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("<= requires at least 2 arguments")
	}
	for i := 0; i < len(args)-1; i++ {
		f1, err1 := toFloat(args[i])
		f2, err2 := toFloat(args[i+1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("<= arguments must be numbers")
		}
		if !(f1 <= f2) {
			return false, nil
		}
	}
	return true, nil
}

func equalNum(args []any) (any, error) {
	if len(args) < 2 {
		return nil, fmt.Errorf("= requires at least 2 arguments")
	}
	for i := 0; i < len(args)-1; i++ {
		f1, err1 := toFloat(args[i])
		f2, err2 := toFloat(args[i+1])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("= arguments must be numbers")
		}
		if f1 != f2 {
			return false, nil
		}
	}
	return true, nil
}

func absFunc(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("abs requires exactly 1 argument, got %d", len(args))
	}
	switch v := args[0].(type) {
	case int:
		if v < 0 {
			return -v, nil
		}
		return v, nil
	case float64:
		return math.Abs(v), nil
	default:
		return nil, fmt.Errorf("abs argument must be a number, got %T", args[0])
	}
}

func powFunc(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("pow requires exactly 2 arguments, got %d", len(args))
	}
	b, err1 := toFloat(args[0])
	e, err2 := toFloat(args[1])
	if err1 != nil || err2 != nil {
		return nil, fmt.Errorf("pow arguments must be numbers")
	}

	res := math.Pow(b, e)

	_, isIntB := args[0].(int)
	_, isIntE := args[1].(int)
	if isIntB && isIntE && e >= 0 && res == math.Floor(res) {
		return int(res), nil
	}
	return res, nil
}

func equalFunc(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("equal? requires exactly 2 arguments, got %d", len(args))
	}
	return reflect.DeepEqual(args[0], args[1]), nil
}

func listFunc(args []any) (any, error) {
	return node.List(args), nil
}

func consFunc(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("cons requires exactly 2 arguments, got %d", len(args))
	}
	elem := args[0]
	l, ok := args[1].(node.List)
	if !ok {
		return nil, fmt.Errorf("second argument to cons must be a list, got %T", args[1])
	}
	newList := make(node.List, 0, len(l)+1)
	newList = append(newList, elem)
	newList = append(newList, l...)
	return newList, nil
}

func carFunc(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("car requires exactly 1 argument, got %d", len(args))
	}
	l, ok := args[0].(node.List)
	if !ok {
		return nil, fmt.Errorf("car argument must be a list, got %T", args[0])
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("car of empty list")
	}
	return l[0], nil
}

func cdrFunc(args []any) (any, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("cdr requires exactly 1 argument, got %d", len(args))
	}
	l, ok := args[0].(node.List)
	if !ok {
		return nil, fmt.Errorf("cdr argument must be a list, got %T", args[0])
	}
	if len(l) == 0 {
		return nil, fmt.Errorf("cdr of empty list")
	}
	return l[1:], nil
}

func mapFunc(args []any) (any, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("map requires exactly 2 arguments, got %d", len(args))
	}

	fn, ok := args[0].(func([]any) (any, error))
	if !ok {
		return nil, fmt.Errorf("first argument to map must be a function, got %T", args[0])
	}

	l, ok := args[1].(node.List)
	if !ok {
		return nil, fmt.Errorf("second argument to map must be a list, got %T", args[1])
	}

	res := make(node.List, len(l))
	for i, item := range l {
		val, err := fn([]any{item})
		if err != nil {
			return nil, fmt.Errorf("error mapping element %d: %w", i, err)
		}
		res[i] = val
	}
	return res, nil
}

func NewGlobalEnv() *Env {
	env := NewEnv(nil)

	// Arithmetic
	env.Set(node.Symbol("+"), add)
	env.Set(node.Symbol("-"), subtract)
	env.Set(node.Symbol("*"), multiply)
	env.Set(node.Symbol("/"), divide)
	env.Set(node.Symbol("mod"), mod)
	env.Set(node.Symbol("incf"), incf)
	env.Set(node.Symbol("decf"), decf)
	env.Set(node.Symbol("abs"), absFunc)
	env.Set(node.Symbol("pow"), powFunc)

	// Constants
	env.Set(node.Symbol("pi"), math.Pi)

	// Comparisons
	env.Set(node.Symbol(">"), greaterThan)
	env.Set(node.Symbol(">="), greaterOrEqual)
	env.Set(node.Symbol("<"), lessThan)
	env.Set(node.Symbol("<="), lessOrEqual)
	env.Set(node.Symbol("="), equalNum)
	env.Set(node.Symbol("equal?"), equalFunc)

	// List operations
	env.Set(node.Symbol("list"), listFunc)
	env.Set(node.Symbol("cons"), consFunc)
	env.Set(node.Symbol("car"), carFunc)
	env.Set(node.Symbol("cdr"), cdrFunc)
	env.Set(node.Symbol("map"), mapFunc)

	return env
}
