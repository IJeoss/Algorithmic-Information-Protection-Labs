package main

import "fmt"

func main(){
	a:=Polynomial{1, 2, 3, 4}
	b:=Polynomial{5, 6, 7, 8}

	fmt.Println("naive:", naiveMul(a, b))
	fmt.Println("karatsuba:", karatsuba(a, b))
	fmt.Println("toom3:", toom3(a, b))
}