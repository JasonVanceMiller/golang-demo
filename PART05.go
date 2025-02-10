//MAPS

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

    //MAP
    myMap := map[string]int{"Alice": 100, "Bob": 200}
    for key, value := range(myMap) {
        fmt.Printf("key: %s, value: %d\n", key, value)
    }
    fmt.Printf("\n")
}
