package main

//полином хранит коэффициенты (степень полинома=размер полинома - 1), макс. степень, под которую выделена память слайса
type Polynomial []float64

const naiveThreshold=32


//функция сложения полиномов за O(max(n, m))
func add(a, b Polynomial) Polynomial{
	maxLen:=max(len(a), len(b))
	result:=make(Polynomial, maxLen)
	copy(result, a)

	for i, bValue:=range b{
		result[i]+=bValue
	}

	return result
}

//функция разности полиномов за O(max(n,m))
func sub(a, b Polynomial)Polynomial{
	maxLen:=max(len(a), len(b))
	result:=make(Polynomial, maxLen)
	copy(result, a)

	for i, bValue:=range b{
		result[i]-=bValue
	}
	return result
}

//умножение на константу за O(n)
func constMul(p Polynomial, k float64)Polynomial{
	if k==0{
		return Polynomial{0}
	}
	result:=make(Polynomial, len(p))

	for i, value:=range p{
		result[i]=value*k
	}
	return result
}
