package main

import "fmt"

func main() {
	var i int
	i = 42

	var j int

	k := 29
	
	fmt.Println("Hello, 世界")
	const greet string = "Greetings earthlings!! \n"

	if k < 29 {
		fmt.Println(k)
	} else if j > 20 {
		fmt.Println(j)
	} else {
		fmt.Println(i + j)
		fmt.Printf(greet)
	}
}
