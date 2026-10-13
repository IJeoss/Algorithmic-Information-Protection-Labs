package main

import "fmt"

//lookupTable хранит таблицу предвычисленных комбинаций бинарных векторов в битово-упакованном виде, элементы каждой строки упаковываются в BitWord в порядке LSB-first
type lookupTable struct{
	combinations int//фактический размер таблицы (2^k)
	words int//количество bitWord в каждой строке (макс. 4 слова)
	data []bitWord
}

func newLookupTable(combinations, words int) lookupTable{
	return lookupTable{combinations, words, make([]bitWord, combinations*words)}
}

func (table lookupTable) getRow(index int) []bitWord{
	if index<0 || index>=table.combinations{
		panic(fmt.Sprintf("lookup table index %d out of range table %d", index, table.combinations))
	}
	start:=index*table.words
	end:=start+table.words
	return table.data[start:end]
}

func buildLookupTable(b matrix, startRow, blockSize int) lookupTable{
	if blockSize<=0 || startRow<0 || startRow+blockSize>b.size{
		panic("invalid parameters for building lookup table")
	}

	combinations:=1<<blockSize//2^blockSize
	table:=newLookupTable(combinations, b.words)
	for bit:=range blockSize{
		row:=b.getRow(startRow+bit)
		offset:=0b1 << bit
		
		for mask:=range offset{
			index:=offset | mask
			copy(table.getRow(index), table.getRow(mask))
			xorRows(table.getRow(index), row)
		}
	}
	return table
}

func getBlockMask(a matrix, row, startCol, blockSize int) int{
	mask:=0
	for bit:=range blockSize{
		if a.getElement(row, startCol+bit){
			mask|=0b1 << bit
		}
	}
	return mask
}

func fourRussiansMultiply(a, b matrix, k int) matrix{
	if a.size!=b.size || a.size<=0 || b.size<=0{
		panic(fmt.Sprintf("matrix sizes do not match for multiplication: a.size=%d, b.size=%d", a.size, b.size))
	}

	if k<=0 || k>a.size{
		panic("invalid block size for four Russians method")
	}

	n:=a.size
	result := newMatrix(n)

	for start:=0; start<n; start+=k{
		blockSize:=min(k, n-start)
		table:=buildLookupTable(b, start, blockSize)

		for i := range n {
			mask := getBlockMask(a, i, start, blockSize)
			xorRows(result.getRow(i), table.getRow(mask))
		}
	}

	return result
}