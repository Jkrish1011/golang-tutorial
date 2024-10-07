package functions

import "errors"

func dividingFunction(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Dividing by zero prohibitted!")
	}
	return a / b, nil
}
