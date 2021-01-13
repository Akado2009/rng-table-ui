package main

import (
	"fmt"
)


func main() {
	fmt.Println(toFixed(10.55, 1))
	fmt.Println(toFixed(10.516, 2))
	fmt.Printf("%+v\n", generateRandomTable(3, 3, 10.0, 11.0))
}
