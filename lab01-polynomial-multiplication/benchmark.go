package main

import "time"

var benchmarkSink float64

//функция измерения алгоритмов умножения
func measureAlgorithm(algorithm func(Polynomial, Polynomial) Polynomial, a, b Polynomial) time.Duration {
	start:=time.Now()

	result:=algorithm(a, b)

	elapsed:=time.Since(start)

	if len(result)>0{
		benchmarkSink=result[len(result)-1]
	}
	return elapsed
}