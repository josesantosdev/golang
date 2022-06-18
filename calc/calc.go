package calc

import (
	"errors"
)

func Dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Não pode dividir por zero")
	}
	return (a / b), nil
}
