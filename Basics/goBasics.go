package main

import (
	"fmt"
	"math/rand"
	"math"
)

func add1(x1 int, x2 int) int{
	return x1+x2
}

/**When two or more consecutive named function parameters share a type,
you can omit the type from all but the last.*/

func add2(x1, x2 int) int {
	return x1+x2
}

func swap2 ( a, b string) (string, string){
	return b , a
}

func modify3 (a, b string) (string, string, string){
	return b, a, a+b
}

func main() {
	fmt.Println("GO BASICS")
	var x int = 590
	y:=10

	if x > 82 {
		fmt.Println(x+y)
	}

	var a [5] int
	a[2] =   38

	fmt.Println(a)

	b:= [] int {5,6,2,4,5}
	b = append(b, 1232)

	fmt.Println(b)

	vertices := make(map[string]int)

	vertices["triangle"] = 2
	vertices["square"] =3

	fmt.Println(vertices)

	delete(vertices,"square")

	fmt.Println(vertices)


	for i:=0; i<3; i++ {
		fmt.Println(i)
	}

	fmt.Println("Random number:: ", rand.Intn(10))

	// Seeding with the same value results in the same random sequence each run.
	// For different numbers, seed with a different value, such as
	// time.Now().UnixNano(), which yields a constantly-changing number.
	//rand.Seed(42)
	
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])

	fmt.Printf("Now you have %g problems.\n", math.Sqrt(707))
	// fmt.Println(math.pi) //
	fmt.Printf("Now you have %g solutions.\n", math.Pi)

	fmt.Println("addition1 :: ",add1(2,3));
	fmt.Println("addition2 :: ",add2(2,3));

	E, P := swap2("Easy" , "Peasy");
	fmt.Println("func returns 2 values :: ", P, E)
	E, P, Z := modify3("Easy" , "Peasy");

	fmt.Println("func returns 2 values :: ", E, P, Z)

}