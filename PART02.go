//FUNCTIONS

package main 

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")

    var helloString string = "Hello"
    worldString := "World"

    fmt.Printf("%s %s\n", helloString, worldString)

    var firstInt int = 111
    secondInt := 222

    fmt.Printf("%d %d\n", firstInt, secondInt)

    fmt.Printf("concat: %s\n", concatStrings(helloString, worldString))

    fmt.Printf("Sum: %d\n", sumInts(firstInt, secondInt))
}

func concatStrings(a string, b string) string {
    return a + b 
}

func sumInts(a int, b int) int {
    return a + b 
}
