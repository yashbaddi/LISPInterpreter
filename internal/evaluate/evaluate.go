package evaluate

import (
	"fmt"

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

func NewGlobalEnv() *Env {
	env := NewEnv(nil)
	env.Set(node.Symbol("+"), func(args []any) (any, error) {
		sum := 0
		for _, arg := range args {
			num, ok := arg.(int)
			if !ok {
				return nil, fmt.Errorf("invalid argument type: expected int, got %T", arg)
			}
			sum += num
		}
		return sum, nil
	})
	env.Set(node.Symbol("-"), func(args []any) (any, error) {
		if len(args) == 0 {
			return 0, fmt.Errorf("subtraction requires at least one argument")
		}
		first, ok := args[0].(int)
		if !ok {
			return nil, fmt.Errorf("invalid argument type: expected int, got %T", args[0])
		}
		if len(args) == 1 {
			return -first, nil
		}
		res := first
		for _, arg := range args[1:] {
			num, ok := arg.(int)
			if !ok {
				return nil, fmt.Errorf("invalid argument type: expected int, got %T", arg)
			}
			res -= num
		}
		return res, nil
	})
	return env
}

func Eval(x any, env *Env) (any, error) {
	switch val := x.(type) {
	case int:
		return val, nil
	case string:
		return val, nil
	case node.Symbol:
		res, ok := env.Get(val)
		if !ok {
			return nil, fmt.Errorf("undefined symbol: %s", val)
		}
		return res, nil
	case node.List:
		if len(val) == 0 {
			return nil, nil
		}
		op, err := Eval(val[0], env)
		if err != nil {
			return nil, err
		}

		var evaluatedArgs []any
		for _, arg := range val[1:] {
			ev, err := Eval(arg, env)
			if err != nil {
				return nil, err
			}
			evaluatedArgs = append(evaluatedArgs, ev)
		}

		switch fn := op.(type) {
		case func([]any) (any, error):
			return fn(evaluatedArgs)
		default:
			return nil, fmt.Errorf("not a function: %v", op)
		}
	default:
		return nil, fmt.Errorf("unknown node type: %T", x)
	}
}
