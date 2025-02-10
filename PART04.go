//LISTS 

package main 

import (
    "fmt"
)

func main() {
    //LISTS 
    myList := []int{10,11,12}
    for index, value := range(myList) {
        fmt.Printf("index: %d, value: %d\n", index, value)
    }
    fmt.Printf("\n")
}
