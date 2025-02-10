//DEFER

package main 

import (
)

func main() {
    //MODERN STYLE
    err, resource1 := acquireResource1()
    if err != nil {
        return
    }
    defer resoruce1.Release()

    err, resource2 := acquireResource2()
    if err != nil {
        return
    }
    defer resoruce2.Release()
     
    err, resource3 := acquireResource3()
    if err != nil {
        return
    }
    defer resoruce3.Release()

    err, resource4 := acquireResource4()
    if err != nil {
        return
    }
    defer resoruce4.Release()

    //DO STUFF
}
