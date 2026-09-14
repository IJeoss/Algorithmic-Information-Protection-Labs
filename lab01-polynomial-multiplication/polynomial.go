package main

//полином хранит коэффициенты (степень полинома=размер полинома - 1)
type Polynomial []float64

const naiveThreshold=32


//функция сложения полиномов за O(max(n, m))
func add(a, b Polynomial) Polynomial{
	maxLen:=max(len(a), len(b))
	result:=make(Polynomial, maxLen)
	copy(result, a)

	for i, b_value:=range b{
		result[i]+=b_value
	}

	return result
}

//функция разности полиномов за O(max(n,m))
func sub(a, b Polynomial)Polynomial{
	maxLen:=max(len(a), len(b))
	result:=make(Polynomial, maxLen)
	copy(result, a)

	for i, b_value:=range b{
		result[i]-=b_value
	}
	return result
}

//умножение на константу за O(max(n,m))
func constMul(p Polynomial, k float64)Polynomial{
	if (k==0){
		return Polynomial{}
	}

	result:=make(Polynomial, len(p))

	for i:=range p{
		result[i]=p[i]*k
	}
	return result
}
