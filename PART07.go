//RESOURCES

package main 

import (
)

func main() {
    //WRONG, AQUIRE BUT NOT RELEASE RESORUCES
    err, resource1 := acquireResource1()
    if err != nil {
        return
    }

    err, resource2 := acquireResource2()
    if err != nil {
        return
    }
     
    err, resource3 := acquireResource3()
    if err != nil {
        return
    }

    err, resource4 := acquireResource4()
    if err != nil {
        return
    }
}
