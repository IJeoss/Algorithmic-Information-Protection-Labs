package main

import "fmt"

const size=2

func main(){
	a:=newMatrix(size)
	b:=newMatrix(size)

	for i:=range size{
		for j:=range size{
			a.SetElement(i, j, (i+j)%2==0)
			b.SetElement(i, j, (i+j)%2==1)
		}
	}
	fmt.Println("Matrix A:")
	fmt.Println()
	fmt.Println(a)
	fmt.Println("Matrix B:")
	fmt.Println()
	fmt.Println(b)

	a.SetElement(23,1, true)

	c:=NaiveMultiply(a, b)
	fmt.Println("Matrix C=A*B:")
	fmt.Println()
	fmt.Println(c)

	d:=SumMatrix(a, b)
	fmt.Println("Matrix D=A+B:")
	fmt.Println()
	fmt.Println(d)

	d.SetElement(1,1,false)
	fmt.Println("Matrix D after setting D[1][1]=0:")
	fmt.Println()
	fmt.Println(d)
}