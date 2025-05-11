package saludos

import "fmt"

func Saludo(name string, age uint8) string {
	message := fmt.Sprintf("Hola %v de %v años", name, age)
	return message
}
