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
func evaluateExpressionWithNumber(exp string, variables map[string]interface{}) int {
	expressions := strings.Split(exp, "=")
	var evaluate interface{}
	var err interface{}
	if len(expressions) > 1 {
		evaluator, _ := govaluate.NewEvaluableExpression(expressions[1])
		evaluate, _ = evaluator.Evaluate(variables)
	} else {
		evaluator, _ := govaluate.NewEvaluableExpression(expressions[0])
		evaluate, _ = evaluator.Evaluate(variables)
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
