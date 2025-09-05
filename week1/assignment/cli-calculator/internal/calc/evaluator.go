package calc

import (
	"errors"
	"regexp"
	"strconv"
)

func Evaluate(expr string) (float64, error) {
	re := regexp.MustCompile(`^([0-9]+\.?[0-9]*)([+\-*/])([0-9]+\.?[0-9]*)$`)
	matches := re.FindStringSubmatch(expr)

	if len(matches) != 4 {
		return 0, errors.New("invalid expression, use format: 10+5")
	}

	a, err1 := strconv.ParseFloat(matches[1], 64)
	op := matches[2]
	b, err2 := strconv.ParseFloat(matches[3], 64)

	if err1 != nil || err2 != nil {
		return 0, errors.New("invalid number")
	}

	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unsupported operator")
	}
}
