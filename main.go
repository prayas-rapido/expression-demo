package main

import (
	"fmt"
	"github.com/maja42/goval"
	"strings"
)

func evaluateKmExpressions(exp string, variables map[string]interface{}) bool {
	expressions := strings.Split(exp, "and")
	evaluator := goval.NewEvaluator()
	for _, expression := range expressions {
		evaluate, err := evaluator.Evaluate(expression, variables, nil)
		if err != nil || evaluate == false {
			return false
		}
	}
	return true
}

func evaluateBaExpressions(exp string, variables map[string]interface{}) interface{} {
	expressions := strings.Split(exp, "=")
	evaluator := goval.NewEvaluator()
	evaluate, err := evaluator.Evaluate(expressions[1], variables, nil)
	fmt.Println(evaluate)
	if err != nil {
		return evaluator
	}
	return evaluator
}
