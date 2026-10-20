package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

func main(){
	size:=flag.Int("size", 8, "matrix dimension: 8, 16, 32, 64, 128 or 256")
	k:=flag.Int("k", 3, "number of rows in a Four Russians block (1..8)")
	flag.Parse()

	if *size < 8 || *size > 256 || (*size&(*size-1)) != 0{
		fmt.Fprintln(os.Stderr, "size must be one of 8, 16, 32, 64, 128, 256")
		os.Exit(2)
	}
	if *k < 1 || *k > 8 || *k > *size{
		fmt.Fprintln(os.Stderr, "k must be between 1 and 8")
		os.Exit(2)
	}

	a, b:=newMatrix(*size), newMatrix(*size)
	for i:=0; i < *size; i++{
		for j:=0; j < *size; j++{
			a.setElement(i, j, (i*7+j*3)%5==0)
			b.setElement(i, j, (i*11+j*13)%7 < 3)
		}
	}

	start:=time.Now()
	result:=fourRussiansMultiply(a, b, *k)
	elapsed:=time.Since(start)

	expected:=naiveMultiply(a, b)
	if !result.isEqual(expected){
		fmt.Fprintln(os.Stderr, "ERROR: results of algorithms differ")
		os.Exit(1)
	}

	fmt.Printf("N=%d, k=%d, Four Russians: %s, correct=true\n", *size, *k, elapsed)
	if *size <= 8{
		fmt.Printf("A:\n%s\nB:\n%s\nA*B:\n%s", a, b, result)
	}
	fmt.Println("For reliable timings use: go test -run '^$' -bench .")
}
