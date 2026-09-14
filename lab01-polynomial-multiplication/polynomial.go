package main

//полином хранит коэффициенты (степень полинома=размер полинома - 1)
type Polynomial []float64


//функция сложения полиномов (O(n))
func add(a, b Polynomial) Polynomial{
	maxLen:=max(len(a), len(b))
	result:=make(Polynomial, maxLen)

	for i, a_value:=range a{
		result[i]=a_value
	}

	for i, b_value:=range b{
		result[i]+=b_value
	}

	return result
}

//функция разности полиномов
func sub(a, b Polynomial)Polynomial{
	maxLen:=max(len(a), len(b))
	result:=make(Polynomial, maxLen)

	for i, a_value:=range a{
		result[i]=a_value
	}

	for i, b_value:=range b{
		result[i]-=b_value
	}
	return result
}