package main

//обычное умножение за O(n*m)
func naiveMul(a, b Polynomial)Polynomial{
	if (len(a)==0 || len(b)==0){//если один из блоков пустой, его произведение с другим блоком ноль
		return Polynomial{}//также пустые блоки могут появляться при рекурсивном разбиении
	}

	result:=make(Polynomial, len(a)+len(b)-1)//len(a)-1 + len(b)-1 + 1 = длина нового (степень+1)

	for i:=range a{
		for j:=range b{
			result[i+j]+=a[i]*b[j]
		}
	}

	return result
}