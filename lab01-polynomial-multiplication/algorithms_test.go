package main

import("math"
	"testing")


const eps=1e-9

//функция сравнения двух многочленов с учетом погрешности eps
func equalPolynomial(a, b Polynomial) bool{
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
func TestPolynomialMultiplication(t *testing.T){
	//небольшие многочлены с заранее известными результатами
	tests:=[]struct{
		name string
		a Polynomial
		b Polynomial
		expected Polynomial
	}{
		{
			name:"constants",
			a:Polynomial{3},
			b:Polynomial{4},
			expected:Polynomial{12},
		},
		{
			name:"linear",
			a:Polynomial{1, 2},
			b:Polynomial{3, 4},
			expected:Polynomial{3, 10, 8},
		},
		{
			name:"different degrees",
			a:Polynomial{1, 2, 3},
			b:Polynomial{4, 5},
			expected:Polynomial{4, 13, 22, 15},
		},
		{
			name:"negative coefficients",
			a:Polynomial{1, -2, 3},
			b:Polynomial{-1, 4},
			expected:Polynomial{-1, 6, -11, 12},
		},
		{
			name:"fractional coefficients",
			a:Polynomial{0.5, 1.5},
			b:Polynomial{2, -0.5},
			expected:Polynomial{1, 2.75, -0.75},
		},
		{
			name:"zero polynomial",
			a:Polynomial{0},
			b:Polynomial{1, 2, 3},
			expected:Polynomial{0, 0, 0},
		},
		{
			name:"recursive case",
			a:Polynomial{1, 0, 0, 0, 1},
			b:Polynomial{1, 0, 0, 1},
			expected:Polynomial{1, 0, 0, 1, 1, 0, 0, 1},
		},
	}

	//все реализации должны давать один и тот же правильный результат
	algorithms:=[]struct{
		name string
		mul func(Polynomial, Polynomial) Polynomial
	}{
		{name:"naive", mul:naiveMul},
		{name:"karatsuba", mul:karatsuba},
		{name:"toom3", mul:toom3},
	}

	for _, test:=range tests{
		for _, algorithm:=range algorithms{
			t.Run(test.name+"/"+algorithm.name, func(t *testing.T){
				got:=algorithm.mul(test.a, test.b)

				if !equalPolynomial(got, test.expected){
					t.Errorf(
						"expected %v, got %v",
						test.expected,
						got,
					)
				}
			})
		}
	}
}