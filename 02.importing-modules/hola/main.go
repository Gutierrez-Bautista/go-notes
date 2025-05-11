package main

import (
	"fmt"

	"example/saludos"
)

func main() {
	ms := saludos.Saludo("Paco", 20)

	fmt.Println(ms)
}