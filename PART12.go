//INTERFACES

package main 

import (
    "fmt"
)

type Dog struct {
    Name string
}
type Cat struct {
    Name string
}

type Animal interface {
    Speak()
}


func main() {
    dog := Dog{"Spot"}
    cat := Cat{"Snowball"}

    Animal(dog).Speak()
    Animal(cat).Speak()
}

func(d Dog)Speak(){
    fmt.Printf("I am a dog named %s\n", d.Name) 
}
func(d Cat)Speak(){
    fmt.Printf("I am a cat named %s\n", d.Name) 
}
