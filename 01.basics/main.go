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
	
	var myArray [3]int8
	myArray[0] = 30
	myArray[1] = 60
	myArray[2] = 120
	fmt.Println(myArray)

	ages := make(map[string]uint8)

	ages["Paco"] = 24
	ages["Carlos"] = 36
	ages["Juan"] = 64

	fmt.Println(ages)

	sonsQty := map[string]uint8{"Paco": 0, "Carlos": 1, "Juan": 4}
	fmt.Println(sonsQty)
}