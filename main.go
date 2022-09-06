package main

import (
	"github.com/antonmedv/expr"
	"strings"
)

func evaluateBooleanExpressions(exp string, variables map[string]interface{}) bool {
	eval, err := expr.Eval(exp, variables)
	if err != nil || eval == false {
		return false
	}
	return true
}

func evaluateExpressionWithNumber(exp string, variables map[string]interface{}) int {
	expressions := strings.Split(exp, "=")
	var evaluate interface{}
	var err interface{}
	if len(expressions) > 1 {
		evaluate, err = expr.Eval(expressions[1], variables)
	} else {
		evaluate, err = expr.Eval(expressions[0], variables)
	}
	if err != nil {
		return -1
	}
	switch v := evaluate.(type) {
	case int:
		return v
	case float64:
		return int(v)
	default:
		return -1
	}
}
