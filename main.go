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

/*
	returns updated BA in case of FLAT and Percentage Offers
	returns discount value given by this coupon in case FIX and Km offers
*/
func evaluateExpressionWithNumber(exp string, variables map[string]interface{}) float64 {
	expression := strings.Split(exp, "=")
	var evaluate interface{}
	if len(expression) > 1 {
		evaluator, _ := govaluate.NewEvaluableExpression(expression[1])
		evaluate, _ = evaluator.Evaluate(variables)
	} else {
		evaluator, _ := govaluate.NewEvaluableExpression(expression[0])
		evaluate, _ = evaluator.Evaluate(variables)
	}

	switch v := evaluate.(type) {
	case float64:
		return v
	default:
		return -1
	}
}
