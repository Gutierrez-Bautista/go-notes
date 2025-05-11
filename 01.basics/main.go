package main

import (
	"fmt"
	"strconv"
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
	
	myArray2 := [3]int8{1, 2}
	fmt.Println(myArray2)

	ages := make(map[string]uint8)

	ages["Paco"] = 24
	ages["Carlos"] = 36
	ages["Juan"] = 64

	fmt.Println(ages)

	sonsQty := map[string]uint8{"Paco": 0, "Carlos": 1, "Juan": 4}
	fmt.Println(sonsQty)
	
	day := "lunes"
	
	switch day {
		case "sábado", "domingo":
			fmt.Println("Fin de semana")
		default:
			fmt.Println("Entre semana")
	}

	myArray3 := [3]uint8{10, 140, 255}

	for index := 0; index < len(myArray3); index++ {
		fmt.Println(myArray3[index])
	}

	myMap := map[string]bool{"key1":true, "key2":false, "key3":false}

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println(saludar("Paco", 30))
}

func saludar (name string, age int) string {
	return "Hola " + name + " - " + strconv.Itoa(age) + " años"
}