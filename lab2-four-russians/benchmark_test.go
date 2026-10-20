package main

import (
	"fmt"
	"testing"
)

var benchmarkResult matrix


func BenchmarkMatrixMultiply(b *testing.B){
	for _, n:=range []int{8, 16, 32, 64, 128, 256}{
		b.Run(fmt.Sprintf("N=%d", n), func(b *testing.B){
			// Матрицы генерируются один раз, за пределами таймера подбенчмарков.
			a, right:=testMatrices(n, int64(n))

			b.Run("reference", func(b *testing.B){
				b.ReportAllocs()
				b.ResetTimer()
				for i:=0; i < b.N; i++{
					benchmarkResult=stupidNaiveMultiply(a, right)
				}
			})

			b.Run("packed_naive", func(b *testing.B){
				b.ReportAllocs()
				b.ResetTimer()
				for i:=0; i < b.N; i++{
					benchmarkResult=naiveMultiply(a, right)
				}
			})

			for k:=1; k <= 8 && k <= n; k++{
				b.Run(fmt.Sprintf("four_russians/k=%d", k), func(b *testing.B){
					b.ReportAllocs()
					b.ResetTimer()
					for i:=0; i < b.N; i++{
						benchmarkResult=fourRussiansMultiply(a, right, k)
					}
				})
			}
		})
	}
}
