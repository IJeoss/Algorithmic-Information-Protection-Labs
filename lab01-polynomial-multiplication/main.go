package main

import("fmt"
	"math/rand"
)

func main(){
	//фиксированный seed !six-seven! позволяет получать одинаковые случайные данные при каждом запуске программы
	rng:=rand.New(rand.NewSource(67))

	//значения n
	ns:=[]int{5, 10, 15, 18}

	for _, n:=range ns{
		//degree=2^n (сдвиг битов на n бит влево, аналог возведения 2 в n-степерь)
		degree:=1<<n

		a:=randomPolynomial(degree, rng)
		b:=randomPolynomial(degree, rng)

		karatsubaTime:=measureAlgorithm(karatsuba, a, b)
		toom3Time:=measureAlgorithm(toom3, a, b)

		fmt.Println("n:", n)
		fmt.Println("degree:", degree)
		fmt.Println("karatsuba:", karatsubaTime)
		fmt.Println("toom3:", toom3Time)
		fmt.Println()
	}
}