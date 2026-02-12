package lib

import "fmt"

func ShowMenu() []string {
	defer fmt.Println("0. Exit")
	conv := []string{
		"Fahrenheit",
		"Reamure",
		"Kelvin",
	}
	for v := range conv {
		fmt.Printf("%d.%s\n", v+1, conv[v])
	}
	return conv
}
