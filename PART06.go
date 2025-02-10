//MAPS CONT

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

    //SAFE ACCESSING MAP
    if val, ok := myMap["Alice"]; ok {
        fmt.Printf("Alice has value %d\n", val)
    } else {
        fmt.Println("Could not find value for Alice")
    }

    if val, ok := myMap["Cameron"]; ok {
        fmt.Printf("Cameron has value %d\n", val)
    } else {
        fmt.Println("Could not find value for Cameron")
    }

    if _, ok := myMap["Diago"]; ok {
        fmt.Println("Key Diago exists")
    } else {
        fmt.Println("Key Diago does not exist")
    }
}
