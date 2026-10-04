package main

import (
	"fmt"
	"strings"
)

const bitsPerWord=64

//Matrix хранит бинарную матрицу в битово-упакованном виде, элементы каждой строки упаковываются в uint64 в порядке LSB-first: столбец 0 соответствует младшему биту, столбец 1 следующему и т.д.
//Поэтому при обычном двоичном выводе uint64 строка визуально выглядит зеркально относительно привычного представления матрицы

//Матрица делится на строки, строки делятся на слова (uint64), слова делятся на биты. Внутри uint64 биты нумеруются от 0 до 63, начиная с младшего бита. Столбцы матрицы нумеруются от 0 до size-1, начиная с левого столбца. Строки матрицы нумеруются от 0 до size-1, начиная с верхней строки
type Matrix struct{
	size int//размерность матрицы
	words int//количество uint64 в каждой строке (макс. 4 слова)
	data []uint64//данные матрицы, хранятся построчно
}

func newMatrix(size int) Matrix{
	wordsPerRow:=(size+bitsPerWord-1)/bitsPerWord
	data:=make([]uint64, size*wordsPerRow)
	return Matrix{size, wordsPerRow, data}
}

func (m Matrix) elementLocation(i, j int) (int, uint64){
	m.checkBounds(i, j)
	bitOffset:=uint(j%bitsPerWord)
	wordIndex:=j/bitsPerWord
	dataIndex:=i*m.words + wordIndex
	mask:=uint64(1) << bitOffset

	return dataIndex, mask
}

func (m Matrix) checkBounds(i, j int){
	if i<0 || i>=m.size || j<0 || j>=m.size{
		panic(fmt.Sprintf("matrix index (%d, %d)  out of range for %dx%d matrix", i, j, m.size, m.size))
	}
}

func (m Matrix) GetElement(i, j int) bool{
	dataIndex, mask:=m.elementLocation(i, j)
	return m.data[dataIndex]&mask!=0
}

func (m *Matrix) SetElement(i, j int, value bool){
	dataIndex, mask:=m.elementLocation(i, j)

	if value{
		m.data[dataIndex]|=mask
	}else{
		m.data[dataIndex]&^=mask
	}
}

func (m Matrix) Row(i int) []uint64{
	if i<0 || i>=m.size{
		panic("matrix index out of range")
	}
	start:=i*m.words
	end:=start+m.words
	return m.data[start:end]
}

func (m Matrix) isEqual(otherMatrix Matrix) bool{
	if m.size!=otherMatrix.size{
		return false
	}
	for i:=range m.data{
		if m.data[i]!=otherMatrix.data[i]{
			return false
		}
	}
	return true
}

//наивное умножение бинарных матриц, O(n^3)
func NaiveMultiply(a, b Matrix) Matrix{
	if a.size!=b.size{
		panic("matrix sizes don't match")
	}
	n:=a.size
	result:=newMatrix(n)

	for i:=range n{
		for j:=range n{
			sum:=false
			for k:=range n{
				if a.GetElement(i, k) && b.GetElement(k, j){
					sum=!sum
				}
			}
			result.SetElement(i, j, sum)
		}
	}
	return result
}

//наивное сложение бинарных матриц, O(n^2)
func SumMatrix(a, b Matrix) Matrix{
	if a.size!=b.size{
		panic("matrix sizes don't match")
	}
	n:=a.size
	result:=newMatrix(n)

	for i:=range a.data{
		result.data[i]=a.data[i]^b.data[i]//XORим слова матриц
	}
	return result
}

func xorRows(dst, src []uint64){
	for i:=range dst{
		dst[i]^=src[i]
	}
}

//метод String() реализует интерфейс fmt.Stringer для удобного вывода матрицы в консоль
func (m Matrix) String() string{
	var str strings.Builder
	for i:=range m.size{
		for j:=range m.size{
			if m.GetElement(i, j){
				str.WriteString("1 ")
			}else{
				str.WriteString("0 ")
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}