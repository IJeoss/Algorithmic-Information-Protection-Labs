package main

import "time"

//функция измерения алгоритмов умножения
func measureAlgorithm(algorithm func(Polynomial, Polynomial) Polynomial, a, b Polynomial) time.Duration {
	start:=time.Now()

	result:=algorithm(a, b)

	elapsed:=time.Since(start)

	
	_=result

	return elapsed
}