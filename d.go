package main

import (
	"fmt"
	"regexp"
	"strings"
)

//imports the package formatted I/O
func main() {
	//While running the file using go run HelloWorld.go the  function gets called
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("hello123 world456 33", -1))
	//prints the digits from String

	words := strings.Fields("vikas landge")
	var j = 1

	//print words with serial number
	for i, word := range words {
		fmt.Print(j)
		fmt.Println(".", word)
		j++
		_ = i

	}
	// print sum of even numbers and odd numbers
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9}
	even := 0
	odd := 0
	for _, num := range nums {
		if num%2 == 0 {
			even += num
		} else {
			odd += num
		}
	}
	fmt.Println("sum of even numbers : ", even)
	fmt.Println("sum of odd numbers : ", odd)
}
