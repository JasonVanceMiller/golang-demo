//GO

package main 

import (
    "fmt"
    "time"
)

func main() {
    //SINGLE THREADED
    count := 0
    for i := 0; i < 10000; i++ {
        add10000(&count)
    }
    fmt.Printf("Single threaded count: %d\n", count)

    //MULTI THREADED
    count = 0
    for i := 0; i < 10000; i++ {
        go add10000(&count)
    }
    time.Sleep(time.Second * 2)
    fmt.Printf("Multi threaded count: %d\n", count)
}

func add10000(a *int){
    for i := 0; i < 10000; i++ {
        (*a)++
    }
}
