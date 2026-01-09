// Tugas Besar MTK Diskrit
// Abdul Kadir Al Attas & Muhammad Ihsan
// Paket 32

package main

import (
	"fmt"
	"math"
)

// ================================
// Soal 5: Relasi Rekurens Iteratif
// an = C1*a(n-1) + C2*a(n-2)
// a0 = 1, a1 = 1
// ================================
func rekuren(C1, C2, N int) int {
	if N == 0 || N == 1 {
		return 1
	}

	a0, a1 := 1, 1
	var an int

	fmt.Println("Proses Perhitungan:")
	fmt.Printf("Suku 0: %d | Suku 1: %d\n", a0, a1)

	for i := 2; i <= N; i++ {
		an = C1*a1 + C2*a0
		fmt.Printf("Suku %d: %d\n", i, an)
		a0 = a1
		a1 = an
	}

	return an
}

// ================================
// Soal 6: Deret Geometri
// S_N = a(1 - r^N) / (1 - r)
// S_inf = a / (1 - r)
// ================================
func deretGeometri(a, r float64, N int) {
	SN := a * (1 - math.Pow(r, float64(N))) / (1 - r)
	SInf := a / (1 - r)
	persen := (SN / SInf) * 100

	fmt.Printf("Sum Berhingga S(%d): %.2f\n", N, SN)
	fmt.Printf("Sum Tak Hingga S(inf): %.2f\n", SInf)
	fmt.Printf("Persentase Kedekatan: %.2f%%\n", persen)
}

// ================================
// Contoh Pemanggilan (Paket 32)
// ================================
func main() {
	// --- Paket 32 ---
	// Soal 5
	C1, C2, N := 3, 2, 9
	fmt.Printf("INPUT: C1=%d, C2=%d, N=%d\n", C1, C2, N)
	hasil := rekuren(C1, C2, N)
	fmt.Printf("HASIL AKHIR Suku ke-%d: %d\n\n", N, hasil)

	// Soal 6
	a, r, Ngeo := 4.0, 0.6, 12
	fmt.Printf("Input Paket: a=%.0f, r=%.1f, N=%d\n", a, r, Ngeo)
	deretGeometri(a, r, Ngeo)
}