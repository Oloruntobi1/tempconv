package main


import (
	"fmt"
	"tempconv/tconv"
)

func main(){


	temp := 45.0

	res := tconv.CelsiusToKelvin(temp)
	fmt.Println(res)
}