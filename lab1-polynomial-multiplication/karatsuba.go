package main

//умножение алгоритмом Карацубы за O(n^log2_(3)))
func karatsuba(a, b Polynomial) Polynomial{
	if (len(a)==0 || len(b)==0){//пустой блок трактуется как нулевой многочлен, произведение с ним 0
		return Polynomial{}
	}

	//базовый случай: на небольших многочленах обычное умножение проще и не нагружает стек
	if (len(a)<=naiveThreshold || len(b)<=naiveThreshold){
		return naiveMul(a,b)
	}

	n:=max(len(a), len(b))
	m:=n/2

	a0:=a[:min(m, len(a))]
	a1:=a[min(m, len(a)):]

	b0:=b[:min(m, len(b))]
	b1:=b[min(m, len(b)):]

	//три рекурсивных умножения Карацубы
	z0:=karatsuba(a0, b0)//a0b0
	z2:=karatsuba(a1, b1)//a1b1
	z1:=karatsuba(add(a0, a1), add(b0, b1))//(ao+a1)*(b0+b1)-a0b0-a1b1=a0b1+a1b0

	//z1=a0b1+a1b0
	z1=sub(z1, z0)
	z1=sub(z1, z2)


	result:=make(Polynomial, len(a)+len(b)-1)
	//z0
	copy(result, z0)

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