package main

import("fmt"
	"math/rand"
	"os"
)

func main(){
	rng:=rand.New(rand.NewSource(67))//фиксированный seed !six-seven! позволяет получать одинаковые случайные данные при каждом запуске программы
	//при этом псевдослучайный генератор детерминирован. стартовое состояние одинаковое, последовательность тоже одинаковая
	file,err:=os.Create("results.csv")
	if err!=nil{
		panic(err)
	}
	defer file.Close()

	fmt.Fprintln(file, "n,degree,karatsuba_ms,toom3_ms")
	//значения n
	for n:=5; n <= 18; n++{
		degree:=1<<n//побитовый сдвиг влево аналог 2^n
		a:=randomPolynomial(degree, rng)
		b:=randomPolynomial(degree, rng)

		karatsubaTime:=measureAlgorithm(karatsuba, a, b)
		toom3Time:=measureAlgorithm(toom3, a, b)

		//перевод duration в миллисекунды
		karatsubaMs:=float64(karatsubaTime)/float64(1e6)
		toom3Ms:=float64(toom3Time)/float64(1e6)

		fmt.Printf("n=%d degree=%d karatsuba=%.3fms toom3=%.3fms\n", n, degree, karatsubaMs, toom3Ms)
		fmt.Fprintf(file, "%d,%d,%.6f,%.6f\n", n, degree, karatsubaMs, toom3Ms)
	}
}