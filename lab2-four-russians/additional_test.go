package main

import (
	"fmt"
	"math/rand"
	"testing"
)

// testMatrices создаёт воспроизводимые бинарные матрицы с плотностью единиц около 50%.
// Генерация вызывается вне измеряемого участка бенчмарков.
func testMatrices(size int, seed int64) (matrix, matrix) {
	rng := rand.New(rand.NewSource(seed))
	a, b := newMatrix(size), newMatrix(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			a.setElement(i, j, rng.Intn(2) == 1)
			b.setElement(i, j, rng.Intn(2) == 1)
		}
	}

	return a, b
}

func TestGetBlockMask(t *testing.T) {
	tests := []struct {
		name      string
		startCol  int
		blockSize int
		setBits   []int
		want      int
	}{
		{"first bit", 0, 1, []int{0}, 0b1},
		{"within a word", 10, 4, []int{10, 12}, 0b0101},
		{"crosses 63/64", 63, 3, []int{63, 65}, 0b101},
		{"crosses two words", 62, 5, []int{62, 64, 66}, 0b10101},
		{"starts at 64", 64, 8, []int{65, 68}, 0b00010010},
		{"last bits", 120, 8, []int{120, 127}, 0b10000001},
		{"block size 16 crosses boundary", 59, 16, []int{59, 63, 64, 74}, 0b1000000000110001},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := newMatrix(128)
			for _, col := range tc.setBits {
				m.setElement(0, col, true)
			}

			got := getBlockMask(m, 0, tc.startCol, tc.blockSize)
			if got != tc.want {
				t.Fatalf("getBlockMask(start=%d, size=%d)=%b; want %b", tc.startCol, tc.blockSize, got, tc.want)
			}
		})
	}
}

func TestGetBlockMaskInvalidParameters(t *testing.T) {
	tests := []struct {
		name      string
		startCol  int
		blockSize int
	}{
		{"negative start", -1, 3},
		{"zero block size", 0, 0},
		{"block exceeds matrix", 127, 2},
		{"block exceeds word length", 0, 65},
	}

	m := newMatrix(128)
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("expected panic for invalid block parameters")
				}
			}()
			_ = getBlockMask(m, 0, tc.startCol, tc.blockSize)
		})
	}
}

func TestFourRussiansAgainstScalarReference(t *testing.T) {
	// Scalar reference — независимый тройной цикл, а не другой вариант XOR строк.
	for _, n := range []int{8, 16, 32, 64, 128, 256} {
		a, b := testMatrices(n, int64(n))
		want := stupidNaiveMultiply(a, b)

		if got := naiveMultiply(a, b); !got.isEqual(want) {
			t.Fatalf("packed naive differs from scalar reference for N=%d", n)
		}

		for k := 1; k <= 8; k++ {
			got := fourRussiansMultiply(a, b, k)
			if !got.isEqual(want) {
				t.Fatalf("Four Russians differs from scalar reference for N=%d, k=%d", n, k)
			}
		}
	}
}

func TestFourRussiansZeroAndIdentity(t *testing.T) {
	for _, n := range []int{8, 64, 128} {
		a, _ := testMatrices(n, 100+int64(n))
		zero := newMatrix(n)
		identity := newMatrix(n)
		for i := 0; i < n; i++ {
			identity.setElement(i, i, true)
		}

		for _, k := range []int{3, 5, 8} {
			if got := fourRussiansMultiply(a, identity, k); !got.isEqual(a) {
				t.Fatalf("A*I != A for N=%d, k=%d", n, k)
			}
			if got := fourRussiansMultiply(identity, a, k); !got.isEqual(a) {
				t.Fatalf("I*A != A for N=%d, k=%d", n, k)
			}
			if got := fourRussiansMultiply(a, zero, k); !got.isEqual(zero) {
				t.Fatalf("A*0 != 0 for N=%d, k=%d", n, k)
			}
			if got := fourRussiansMultiply(zero, a, k); !got.isEqual(zero) {
				t.Fatalf("0*A != 0 for N=%d, k=%d", n, k)
			}
		}
	}
}

func TestFourRussiansInvalidK(t *testing.T) {
	tests := []struct {
		n int
		k int
	}{
		{8, -1},
		{8, 0},
		{8, 9},   // k превышает размер матрицы
		{32, 17}, // k превышает ограничение lookup-таблицы (16)
	}

	for _, tc := range tests {
		t.Run(fmt.Sprintf("N=%d/k=%d", tc.n, tc.k), func(t *testing.T) {
			a, b := newMatrix(tc.n), newMatrix(tc.n)
			defer func() {
				if recover() == nil {
					t.Errorf("expected panic for N=%d, k=%d", tc.n, tc.k)
				}
			}()
			_ = fourRussiansMultiply(a, b, tc.k)
		})
	}
}

func TestFourRussiansDifferentSizes(t *testing.T) {
	a, b := newMatrix(8), newMatrix(16)
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for matrices of different sizes")
		}
	}()
	_ = fourRussiansMultiply(a, b, 3)
}
