//STRUCTS & INHERITANCE CONT

package main 

import (
    "fmt"
)

type Parent struct {
    Name string
}
type Child struct {
    Parent
    Name string
}


func main() {
    child := Child{}
    child.Name = "Billy Jr"
    child.Parent.Name = "Billy Sr"

    child.SayName()
}

func(p *Parent)SayName(){
    fmt.Printf("My Name is %s", p.Name) 
}
func(c *Child)SayName(){
    fmt.Printf("My Name is %s", c.Name) 
}
