//RESOURCES CONT

package main 

import (
)

func main() {
    //C STYLE, ERROR PRONE
    err, resource1 := acquireResource1()
    if err != nil {
        return
    }

    err, resource2 := acquireResource2()
    if err != nil {
        resource1.Release()
        return
    }
     
    err, resource3 := acquireResource3()
    if err != nil {
        resource1.Release()
        resource2.Release()
        return
    }

    err, resource4 := acquireResource4()
    if err != nil {
        resource1.Release()
        resource2.Release()
        resource3.Release()
        return
    }
    
    //DO STUFF

    
    resource1.Release()
    resource2.Release()
    resource3.Release()
    resource4.Release()
}
