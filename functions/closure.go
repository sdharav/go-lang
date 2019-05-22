package main

import "fmt"

func main(){
	x := 0
    increment := func() int{
        x += 1
        return x
    }
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
}
