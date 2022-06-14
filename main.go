package main

import (
	"github.com/Knetic/govaluate"
	"strings"
)

func evaluateBooleanExpressions(exp string, variables map[string]interface{}) bool {
	expressions := strings.Split(exp, "and")
	for _, item := range expressions {
		expression, _ := govaluate.NewEvaluableExpression(item)
		eval, err := expression.Evaluate(variables)
		if err != nil || eval == false {
			return false
		}
	}
	return true
}

func evaluateExpressionWithNumber(exp string, variables map[string]interface{}) int {
	expression := strings.Split(exp, "=")
	evaluator, err := govaluate.NewEvaluableExpression(expression[1])
	evaluate, err := evaluator.Evaluate(variables)

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
