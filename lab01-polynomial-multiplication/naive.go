package main

//обычное умножение за O(n^2)
func naiveMul(a, b Polynomial)Polynomial{
	if (len(a)==0 || len(b)==0){
		return Polynomial{}//случай умножения на нуль
	}

	result:=make(Polynomial, len(a)+len(b)-1)//len(a)-1 + len(b)-1 + 1 = длина нового (степень+1)

	for i:=range a{
		for j:=range b{
			result[i+j]+=a[i]*b[j]
		}
	}

	return result
}