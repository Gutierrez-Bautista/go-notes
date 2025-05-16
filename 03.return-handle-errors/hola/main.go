package main

import (
	"fmt"
	"log"

	"example/saludos"
)

func main() {
	log.SetPrefix("Saludo: ")
	log.SetFlags(0)

	ms, err := saludos.Saludo("", 20)

	if err != nil {
		log.Fatal("No name entered")
	}

	fmt.Println(ms)
}