package main

import (
	"fmt"
	"regexp"
)

//imports the package formatted I/O
func main() {
	//While running the file using go run HelloWorld.go the  function gets called
	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("hello123 world456 33", -1))
	//prints the digits from String

}
