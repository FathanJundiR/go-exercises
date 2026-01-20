package main

import (
	"fmt"
	"strconv"
	"strings"
)

var emptyArray = [10]int{}

func prettyPrint(data []string){
	fmt.Printf("Length: %v\n", len(data))
	fmt.Printf("Contents: [%v]\n", strings.Join(data, ","))
}

func typePrint(data any) {
	fmt.Printf("Type: %T\n", data)
}

func number1(data [10]int) {
	fmt.Printf("Length: %v\n", len(data))
	fmt.Printf("Type: %T\n", data)
	fmt.Printf("Contents: %v\n", data)
}

func number1Slice(data []int) {
	fmt.Printf("Length: %v\n", len(data))
	fmt.Printf("Type: %T\n", data)
	fmt.Printf("Contents: %v\n", data)
}

func number2() {
	var strRand = [4]string{}
	for i, _ := range strRand{
		strRand[i] = strconv.Itoa(i)
	}
	typePrint(strRand)
	prettyPrint(strRand[:])
}

func number5() {
	var games = [10]string{}
	games[0] = "Donkey Kong"
	typePrint(games)
	prettyPrint(games[:])
}

func number6() {
	var movies = [4]string{
		"JJK",
		"AOT",
		"KNY",
		"Konosuba",
	}
	typePrint(movies)
	prettyPrint(movies[:])
	fmt.Println("Second element: ", movies[1])
}
func main() {
	// SKIP NUMBER 3 & 4
	fmt.Println("==== NUMBER 1 ====")
	number1(emptyArray)
	number1Slice(emptyArray[:])
	fmt.Println("==== NUMBER 2 ====")
	number2()
	fmt.Println("==== NUMBER 5 ====")
	number5()
	fmt.Println("==== NUMBER 6 ====")
	number6()
}
