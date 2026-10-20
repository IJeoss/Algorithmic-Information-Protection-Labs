package main

import "testing"

func TestGetSetElement(t *testing.T) {
	m := newMatrix(128)

	m.setElement(0, 0, true)
	m.setElement(0, 63, true)
	m.setElement(0, 64, true)
	m.setElement(127, 127, true)

	if !m.getElement(0, 0) {
		t.Error("expected m[0][0]=1")
	}
	if !m.getElement(0, 63) {
		t.Error("expected m[0][63]=1")
	}
	if !m.getElement(0, 64) {
		t.Error("expected m[0][64]=1")
	}
	if !m.getElement(127, 127) {
		t.Error("expected m[127][127]=1")
	}
	if m.getElement(1, 1) {
		t.Error("expected m[1][1]=0")
	}

	m.setElement(0, 64, false)
	if m.getElement(0, 64) {
		t.Error("expected m[0][64]=0 after reset")
	}
}

func TestSumMatrix(t *testing.T) {
	for _, n := range []int{8, 64, 128, 256} {
		a := newMatrix(n)
		b := newMatrix(n)

		for i := range n {
			for j := range n {
				a.setElement(i, j, (i+j)%3 == 0)
				b.setElement(i, j, (i*3+j)%5 < 2)
			}
		}

		// В Z_2 каждый элемент XOR сам с собой даёт 0.
		if got := sumMatrix(a, a); !got.isEqual(newMatrix(n)) {
			t.Errorf("A + A != 0 for N=%d", n)
		}

		// Проверяем содержимое всех битов, в том числе на границах uint64.
		got := sumMatrix(a, b)
		for i := range n {
			for j := range n {
				want := a.getElement(i, j) != b.getElement(i, j)
				if got.getElement(i, j) != want {
					t.Fatalf("wrong sum for N=%d, position=(%d,%d)", n, i, j)
				}
			}
		}
	}
}
