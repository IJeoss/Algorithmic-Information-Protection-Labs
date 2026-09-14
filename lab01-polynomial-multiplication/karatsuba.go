package main

//умножение алгоритмом Карацубы за O(n^log2_(3)))
func karatsuba(a, b Polynomial) Polynomial{
	if (len(a)==0 || len(b)==0){
		return Polynomial{}
	}

	//на небольших многочленах обычное умножение проще и не нагружает стек
	if (len(a)<=2 || len(b)<=2){
		return naiveMul(a,b)
	}

	n:=max(len(a), len(b))
	m:=n/2

	a0:=a[:min(m, len(a))]
	a1:=a[min(m, len(a)):]

	b0:=b[:min(m, len(b))]
	b1:=b[min(m, len(b)):]

	//три рекурсивных умножения Карацубы
	z0:=karatsuba(a0, b0)
	z2:=karatsuba(a1, b1)

	z1:=karatsuba(add(a0, a1), add(b0, b1))

	//z1=a0*b1+a1*b0
	z1=sub(z1, z0)
	z1=sub(z1, z2)

	result:=make(Polynomial, len(a)+len(b)-1)

	//z0
	for i:=range z0{
		result[i]+=z0[i]
	}

	//z1 * x^m
	for i:=range z1{
		if i+m<len(result){
			result[i+m]+=z1[i]
		}
	}

	//z2 * x^(2m)
	for i:=range z2{
		if i+2*m<len(result){
			result[i+2*m]+=z2[i]
		}
	}

	return result
}