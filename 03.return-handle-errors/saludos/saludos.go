package saludos

import (
	"errors"
	"fmt"
)

func Saludo(name string, age uint8) (string, error) {
	if name == "" {
		return "", errors.New("no name")
	}
	message := fmt.Sprintf("Hola %v de %v años", name, age)
	return message, nil
}
