// Package wordy translates a sentence into a math problem
package wordy

import (
	"regexp"
	"strconv"
)

type operator func(op1, op2 int) int

var operators = map[string]operator{
	"plus":          plus,
	"minus":         minus,
	"multiplied by": mult,
	"divided by":    div,
}

func plus(op1, op2 int) int  { return op1 + op2 }
func minus(op1, op2 int) int { return op1 - op2 }
func mult(op1, op2 int) int  { return op1 * op2 }
func div(op1, op2 int) int   { return op1 / op2 }

// Answer finds the answer to a math problem expressed in natural language
func Answer(q string) (a int, ok bool) {
	question := getQuestion(q)
	hasMore := true
	for hasMore {
		r := regexp.MustCompile("^(-?\\d+)\\s+([a-z )]+)\\s+(-?\\d+)?(.*)")
		matches := r.FindAllStringSubmatch(question, 5)
		if matches == nil {
			return 0, false
		}
		partial, ok := execute(matches[0][1], matches[0][2], matches[0][3])
		question = partial + matches[0][4]
		if !ok || matches[0][4] == "" {
			hasMore = false
			continue
		}
	}
	res, err := strconv.Atoi(question)
	return res, err == nil
}

func execute(op1, op, op2 string) (result string, ok bool) {
	o1, err1 := strconv.Atoi(op1)
	o2, err2 := strconv.Atoi(op2)
	if err1 != nil || err2 != nil {
		return "", false
	}

	for name, function := range operators {
		if op == name {
			return strconv.Itoa(function(o1, o2)), true
		}
	}
	return "", false
}

func getQuestion(q string) string {
	r := regexp.MustCompile("^What is (.*)\\?$")
	matches := r.FindAllStringSubmatch(q, 1)
	if matches == nil {
		return ""
	}
	return matches[0][1]
}
