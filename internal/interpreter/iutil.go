package interpreter

import (
	"GLox/internal/loxerror"
	"GLox/internal/scanner/token"
)

// isTruth Lox中规定nil和false为"假"，其余都为真
func isTruth(any interface{}) bool {
	if any == nil {
		return false
	}
	// bool类型直接返回自身
	if val, ok := any.(bool); ok {
		return val
	}
	// 其余一律返回true (nil除外)

	return true
}

func doPlus(operator *token.Token, left, right interface{}) (interface{}, error) {
	_, ok1 := left.(float64)
	_, ok2 := right.(float64)
	if ok1 && ok2 {
		return left.(float64) + right.(float64), nil
	}

	_, ok1 = left.(string)
	_, ok2 = right.(string)
	if ok1 && ok2 {
		return left.(string) + right.(string), nil
	}

	//panic(loxerror.NewRuntimeError(operator, "Operands must be two numbers or two strings."))
	return nil, loxerror.NewRuntimeError(operator, "Operands must be two numbers or two strings.")
}

func isEqual(left, right interface{}) bool {
	if left == nil && right == nil {
		// 不能直接比较 nil == nil
		return true
	}
	// 基础类型也不能和nil直接比较
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}

	return left == right
}

func checkNumberOperands(operator *token.Token, operands ...interface{}) error {
	for _, operand := range operands {
		if _, ok := operand.(float64); !ok {
			//panic(loxerror.NewRuntimeError(operator, "Operand must be a number."))
			return loxerror.NewRuntimeError(operator, "Operand must be a number.")
		}
	}

	return nil
}
