//GO WITH CHANNELS

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

    //MULTI THREADED WITH CHANNEL
    count = 0
    chStart := make(chan(struct{}))
    chDone  := make(chan(struct{}))

    for i := 0; i < 10000; i++ {
        go add10000WithChannel(&count, chStart, chDone)
    }
    //Unlock threads one at a time 
    for i := 0; i < 10000; i++ {
        chStart <- struct{}{}
        <- chDone
    }
    fmt.Printf("Multi threaded with channel count: %d\n", count)
}

func add10000(a *int){
    for i := 0; i < 10000; i++ {
        (*a)++
    }
}
func add10000WithChannel(a *int, chStart chan(struct{}), chDone chan(struct{})){
    <- chStart
    for i := 0; i < 10000; i++ {
        (*a)++
    }
    chDone <- struct{}{}
}
