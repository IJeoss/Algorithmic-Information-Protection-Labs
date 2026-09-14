package main

import "fmt"

func main(){
	a:=Polynomial{2,3}
	b:=Polynomial{4,5}
	c:=naiveMul(a,b)

	fmt.Println(c)
}