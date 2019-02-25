package main

import (
	"calc/sc"
	"calc/simple"
	"fmt"
)

func main() {

	//k := calc.Add(20, 20)
	//i := calc.simple.Add(20, 20)

	//fmt.Printf("Add from main calc  %d ", k)
	fmt.Printf("Add from simple   %d ", simple.Add(20, 20))
	fmt.Printf("Sub from simple  %d ", simple.Sub(20, 20))
	fmt.Printf("Div from simple  %d ", simple.Sub(20, 20))
	fmt.Printf("Mod from simple  %d ", simple.Sub(20, 20))
	fmt.Printf("Sin from Sc  %d ", sc.Sin(30, 30))

}
