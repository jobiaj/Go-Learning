package main

import "fmt"

func main() {
   var a, b, c, d = 3, 4, "foo", 3.34
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
   fmt.Printf("d is of type %T\n", d)
}