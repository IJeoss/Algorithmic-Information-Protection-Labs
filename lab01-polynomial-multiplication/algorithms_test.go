package main

import("math"
	"testing")

//функция сравнения двух многочленов с учетом погрешности eps
func equalPolynomial(a, b Polynomial, eps float64) bool{
	if (len(a)!=len(b)){
		return false
	}

	for i:=range a{
		if math.Abs(a[i]-b[i])>eps{
			return false
		}
	}
	return true
}

func TestNaiveMul(t *testing.T){
	a:=Polynomial{1,2,3}
	b:=Polynomial{4,5}

	expected:=Polynomial{4, 13, 22, 15}
	got:=naiveMul(a,b)
	if !equalPolynomial(expected, got, 1e-9){
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestKaratsuba(t *testing.T){
	a:=Polynomial{1, 2, 3, 4}
	b:=Polynomial{5, 6, 7, 8}

	expected:=naiveMul(a, b)
	got:=karatsuba(a, b)

	if !equalPolynomial(got, expected, 1e-9){
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestToom3(t *testing.T){
	a:=Polynomial{1, 2, 3, 4}
	b:=Polynomial{5, 6, 7, 8}

	expected:=naiveMul(a, b)
	got:=toom3(a, b)

	if !equalPolynomial(got, expected, 1e-9){
		t.Errorf("expected %v, got %v", expected, got)
	}
}