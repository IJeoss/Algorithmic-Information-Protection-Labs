package main

import "math/rand"

//создает рандомный многочлен заданной степени
//rng подается снаружи
func randomPolynomial(degree int, rng *rand.Rand) Polynomial{
	p:=make(Polynomial, degree+1)
	for i:=range p{
		p[i]=rng.Float64()*20-10
	}

	return p
}