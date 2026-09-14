package main

//функция разделения полинома p на 3 части длины примерно m
func split3(p Polynomial, m int) (Polynomial, Polynomial, Polynomial){
	end0:=min(m, len(p))
	end1:=min(2*m, len(p))

	p0:=p[:end0]
	p1:=p[end0:end1]
	p2:=p[end1:]

	return p0, p1, p2
}

//алгоритм Тоома-Кука при k=3
func toom3(a, b Polynomial) Polynomial{
	if (len(a)==0 || len(b)==0){
		return Polynomial{}
	}

	//на маленьких входах можно заменить обычным умножением
	if (len(a)<=3 || len(b)<=3){
		return naiveMul(a, b)
	}

	n:=max(len(a), len(b))
	m:=(n+2)/3//n/k=(n+k-1)/k

	a0, a1, a2:=split3(a, m)
	b0, b1, b2:=split3(b, m)

	//A(0)=A0, B(0)=B0
	aAt0:=a0
	bAt0:=b0

	//A(1)=A0+A1+A2
	aAt1:=add(add(a0, a1), a2)
	bAt1:=add(add(b0, b1), b2)

	//A(-1)=A0-A1+A2
	aAtMinus1:=add(sub(a0, a1), a2)
	bAtMinus1:=add(sub(b0, b1), b2)

	//A(2)=A0+2*A1+4*A2
	aAt2:=add(add(a0, constMul(a1, 2)), constMul(a2, 4))
	bAt2:=add(add(b0, constMul(b1, 2)), constMul(b2, 4))

	//в точке "бесконечность" берём старший блок
	aAtInf:=a2
	bAtInf:=b2

	v0:=toom3(aAt0, bAt0)
	v1:=toom3(aAt1, bAt1)
	vMinus1:=toom3(aAtMinus1, bAtMinus1)
	v2:=toom3(aAt2, bAt2)
	vInf:=toom3(aAtInf, bAtInf)

	c0:=v0
	c4:=vInf

	//s1=C1+C2+C3
	s1:=sub(sub(v1, c0), c4)

	//sMinus1=-C1+C2-C3
	sMinus1:=sub(sub(vMinus1, c0), c4)
	c2:=constMul(add(s1, sMinus1), 0.5)

	//t=C1+C3
	t:=constMul(sub(s1, sMinus1), 0.5)

	//убираем из C(2) уже известные C0 и 16*C4
	temp:=sub(sub(v2, c0), constMul(c4, 16))

	//u=C1+2*C2+4*C3
	u:=constMul(temp, 0.5)
	c3:=constMul(sub(sub(u, t),constMul(c2, 2)), 1.0/3.0)
	c1:=sub(t, c3)

	result:=make(Polynomial, len(a)+len(b)-1)

	for i:=range c0{
		if i<len(result){
			result[i]+=c0[i]
		}
	}

	for i:=range c1{
		if i+m<len(result){
			result[i+m]+=c1[i]
		}
	}

	for i:=range c2{
		if i+2*m<len(result){
			result[i+2*m]+=c2[i]
		}
	}

	for i:=range c3{
		if i+3*m<len(result){
			result[i+3*m]+=c3[i]
		}
	}

	for i:=range c4{
		if i+4*m<len(result){
			result[i+4*m]+=c4[i]
		}
	}

	return result
}