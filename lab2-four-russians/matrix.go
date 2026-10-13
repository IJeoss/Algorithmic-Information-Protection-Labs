package main

import (
	"fmt"
	"strings"
	"slices"
)

const bitsPerWord=64

type bitWord=uint64

//следует отметить, что в силу моей реализации, мы имеем всюду O(N^x/bitsPerWord)

//matrix хранит бинарную матрицу в битово-упакованном виде, элементы каждой строки упаковываются в uint64 в порядке LSB-first: столбец 0 соответствует младшему биту, столбец 1 следующему и т.д.
//поэтому при обычном двоичном выводе uint64 строка визуально выглядит зеркально относительно привычного представления матрицы

//матрица делится на строки, строки делятся на слова (uint64), слова делятся на биты. Внутри uint64 биты нумеруются от 0 до 63, начиная с младшего бита. Столбцы матрицы нумеруются от 0 до size-1, начиная с левого столбца. Строки матрицы нумеруются от 0 до size-1, начиная с верхней строки
type matrix struct{
	size int//размерность матрицы
	words int//количество uint64 в каждой строке (макс. 4 слова)
	data []bitWord
}

func newMatrix(size int) matrix{
	if size<=0{
		panic("matrix size must be positive")
	}
	wordsPerRow:=(size+bitsPerWord-1)/bitsPerWord//целое количество слов в строке, округление вверх
	data:=make([]bitWord, size*wordsPerRow)
	return matrix{size, wordsPerRow, data}
}

func (m matrix) checkBounds(i, j int){
	if i<0 || i>=m.size || j<0 || j>=m.size{
		panic(fmt.Sprintf("matrix index (%d, %d)  out of range for %dx%d matrix", i, j, m.size, m.size))
	}
}

func (m matrix) elementLocation(i, j int) (dataIndex int, mask bitWord){
	m.checkBounds(i, j)
	bitOffset:=j%bitsPerWord
	wordIndex:=j/bitsPerWord
	dataIndex=i*m.words + wordIndex
	mask=0b1 << bitOffset

	return dataIndex, mask
}

func (m matrix) getElement(i, j int) bool{
	dataIndex, mask:=m.elementLocation(i, j)
	return (m.data[dataIndex] & mask) != 0
}

func (m *matrix) setElement(i, j int, value bool){
	dataIndex, mask:=m.elementLocation(i, j)

	if value{
		m.data[dataIndex]|=mask
	}else{
		m.data[dataIndex]&^=mask//&^= - побитовое AND NOT, очищает бит
	}
}

func (m matrix) getRow(index int) []bitWord{
	if index<0 || index>=m.size{
		panic(fmt.Sprintf("matrix index %d out of range matrix %d", index, m.size))
	}
	start:=index*m.words
	end:=start+m.words
	return m.data[start:end]
}

func xorRows(dst, src []bitWord){
	for i:=range dst{
		dst[i]^=src[i]
	}
}

func (m matrix) isEqual(otherMatrix matrix) bool{
	return m.size==otherMatrix.size && slices.Equal(m.data, otherMatrix.data)
}

//наивное умножение бинарных матриц, примерно O(n^3/bitsPerWord) операций над машинными словами
func naiveMultiply(a, b matrix) matrix{
	if a.size!=b.size{
		panic("matrix sizes don't match")
	}
	n:=a.size
	result:=newMatrix(n)

	for i:=range n{
		resultRow:=result.getRow(i)
		for j:=range n{
			if a.getElement(i,j){
				xorRows(resultRow, b.getRow(j))
			}
		}
	}
	return result
}

//глупое навивное умножение бинарных матриц, O(n^3)
func stupidNaiveMultiply(a, b matrix) matrix{
	if a.size!=b.size{
		panic("matrix sizes don't match")
	}

	n:=a.size
	result:=newMatrix(n)

	for i:=range n{
		for j:=range n{
			sum:=false
			for k:=range n{
				if a.getElement(i,k) && b.getElement(k,j){
					sum=!sum
				}
			}
			result.setElement(i, j, sum)
		}
	}
	return result
}

//наивное сложение бинарных матриц, O(n^2)
func sumMatrix(a, b matrix) matrix{
	if a.size!=b.size{
		panic("matrix sizes don't match")
	}
	result:=newMatrix(a.size)

	for i:=range a.data{
		result.data[i]=a.data[i]^b.data[i]//XORим слова матриц
	}
	return result
}

//метод String() реализует интерфейс fmt.Stringer для удобного вывода матрицы в консоль
func (m matrix) String() string{
	var str strings.Builder
	for i:=range m.size{
		for j:=range m.size{
			if m.getElement(i, j){
				str.WriteString("1 ")
			}else{
				str.WriteString("0 ")
			}
		}
		str.WriteString("\n")
	}
	return str.String()
}