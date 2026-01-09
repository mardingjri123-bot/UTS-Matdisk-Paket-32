// Tugas Besar MTK Diskrit
// Abdul Kadir Al Attas & Muhammad Ihsan
// Paket 32

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func buatMatriks(ordo int) [][]int {
	rand.Seed(time.Now().UnixNano())
	data := make([][]int, ordo)

	for i := 0; i < ordo; i++ {
		baris := make([]int, ordo)
		for j := 0; j < ordo; j++ {
			baris[j] = rand.Intn(20) + 1
		}
		data[i] = baris
	}
	return data
}

func kaliMatriks(X, Y [][]int) [][]int {
	n := len(X)
	hasil := make([][]int, n)

	for i := 0; i < n; i++ {
		hasil[i] = make([]int, n)
		for j := 0; j < n; j++ {
			jumlah := 0
			for k := 0; k < n; k++ {
				jumlah += X[i][k] * Y[k][j]
			}
			hasil[i][j] = jumlah
		}
	}
	return hasil
}

func hitungTrace(M [][]int) int {
	total := 0
	for i := 0; i < len(M); i++ {
		total += M[i][i]
	}
	return total
}

func soalTiga(n int) {
	A := buatMatriks(n)
	B := buatMatriks(n)

	fmt.Println("Matrix A:", A)
	fmt.Println("Matrix B:", B)

	R := kaliMatriks(A, B)
	fmt.Println("Hasil Matriks R:", R)
	fmt.Println("Trace:", hitungTrace(R))
}

func soalEmpat(n int) {
	M := buatMatriks(n)
	fmt.Println("Matrix M (Generated):", M)

	// tukar baris pertama dan terakhir
	M[0], M[n-1] = M[n-1], M[0]

	max := M[0][0]
	r, c := 0, 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if M[i][j] > max {
				max = M[i][j]
				r = i
				c = j
			}
		}
	}

	fmt.Println("Matrix Setelah Ditukar:", M)
	fmt.Printf("Nilai Maksimum: %d di Posisi (%d,%d)\n", max, r, c)
}