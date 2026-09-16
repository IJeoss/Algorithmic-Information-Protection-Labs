package main

import "time"

//хранит часть результата алгоритма, чтобы результат вычисления не был полностью неиспользуемым при измерении производительности
var benchmarkSink float64

//функция измерения времени выполнения алгоритмов умножения
func measureAlgorithm(algorithm func(Polynomial, Polynomial) Polynomial, a, b Polynomial) time.Duration{
	start:=time.Now()

	result:=algorithm(a, b)

	elapsed:=time.Since(start)

	if len(result)>0{
		benchmarkSink=result[len(result)-1]
	}
	return elapsed
}