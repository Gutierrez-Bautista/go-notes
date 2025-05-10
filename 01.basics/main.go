package main

import (
	"fmt"
)

func main() {
	const myString string = "This is a string"
	fmt.Println(myString)

	myInt := 10

	if myInt == 0 {
    fmt.Println("myInt es exactamente 0")
  } else if myInt < 0 {
    fmt.Println("myInt es negativo")
  } else {
    fmt.Println("myInt es positivo")
  }

	var myCondition bool = myInt >= 0 && myInt < 130

  if myCondition {
		fmt.Println("myInt es una edad válida")
	}
}