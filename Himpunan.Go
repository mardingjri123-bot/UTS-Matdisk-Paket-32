// Tugas Besar MTK Diskrit
// Abdul Kadir Al Attas & Muhammad Ihsan
// Paket 32

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func symmetricDifference(A, B []int) []int {
	result := []int{}
	for _, a := range A {
		if !contains(B, a) && !contains(result, a) {
			result = append(result, a)
		}
	}
	for _, b := range B {
		if !contains(A, b) && !contains(result, b) {
			result = append(result, b)
		}
	}
	return result
}

func difference(A, C []int) []int {
	result := []int{}
	for _, a := range A {
		if !contains(C, a) && !contains(result, a) {
			result = append(result, a)
		}
	}
	return result
}

func intersection(A, C []int) []int {
	result := []int{}
	for _, a := range A {
		if contains(C, a) && !contains(result, a) {
			result = append(result, a)
		}
	}
	return result
}

func union(A, B []int) []int {
	result := []int{}
	for _, a := range A {
		if !contains(result, a) {
			result = append(result, a)
		}
	}
	for _, b := range B {
		if !contains(result, b) {
			result = append(result, b)
		}
	}
	return result
}

func soalSatu(N int) {
	rand.Seed(time.Now().UnixNano())

	A := []int{5, 12, 18, 25, 40}
	B := []int{12, 20, 25, 33, 50}
	C := []int{5, 20, 33}

	fmt.Println("Generating Sets... (N =", N, ")")
	fmt.Println("A:", A, "| B:", B, "| C:", C)

	R := union(
		difference(symmetricDifference(A, B), C),
		intersection(A, C),
	)

	fmt.Println("Hasil Operasi R:", R)

	filter := []int{}
	batas := N / 4
	for _, x := range R {
		if x%2 == 0 && x < batas {
			filter = append(filter, x)
		}
	}

	fmt.Println("Hasil Filter (Genap <", batas, "):", filter)
	fmt.Println("Total Elemen:", len(filter))
}

func soalDua(M, K int) {
	S := make([]int, M)
	for i := 0; i < M; i++ {
		S[i] = (i + 1) * 2
	}

	fmt.Println("\nSet S:", S, "| Target K:", K)
	fmt.Println("Subset 2-Elemen (Sum >", K, "):")

	count := 0
	for i := 0; i < len(S); i++ {
		for j := i + 1; j < len(S); j++ {
			if S[i]+S[j] > K {
				count++
				fmt.Printf("%d. {%d, %d} (Sum=%d)\n",
					count, S[i], S[j], S[i]+S[j])
			}
		}
	}
	fmt.Println("Total:", count, "Pasangan")
}