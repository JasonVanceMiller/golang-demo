//MULTIPLE RETURN

package main 

import (
    "fmt"
)

func main() {
    fmt.Println("Hello World")

    val, err :=  divideInts(10, 2)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(val)
    }

    val, err =  divideInts(10, 0)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(val)
    }
}

func divideInts(a int, b int) (int, error) {
    if (b == 0) {
        return 0, fmt.Errorf("Divide by zero error for %d / %d", a, b)
    }
    return a / b, nil
}
