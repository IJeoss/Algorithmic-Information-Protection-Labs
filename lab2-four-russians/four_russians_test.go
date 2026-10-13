package main

import "testing"

func TestFourRussiansMultiply(t *testing.T) {
	sizes := []int{8, 16, 32, 64, 128, 256}
	blockSizes := []int{1, 2, 3, 5, 8}

	for _, n := range sizes {
		a := newMatrix(n)
		b := newMatrix(n)

		for i := range n {
			for j := range n {
				a.setElement(i, j, (i*7+j*3)%5 == 0)
				b.setElement(i, j, (i*11+j*13)%7 < 3)
			}
		}

		expected := naiveMultiply(a, b)

		for _, k := range blockSizes {
			result := fourRussiansMultiply(a, b, k)

			if !result.isEqual(expected) {
				t.Fatalf(
					"incorrect result for n=%d, k=%d",
					n, k,
				)
			}
		}
	}
}