package evaluate

import (
	"fmt"

	"github.com/yashbaddi/golisp/internal/node"
)

func Eval(x any, env *Env) (any, error) {
	switch val := x.(type) {
	case int, float64, bool, string:
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
		if sym, ok := val[0].(node.Symbol); ok && isSpecialForm(sym) {
			return evalSpecialForm(sym, val[1:], env)
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
