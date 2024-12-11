package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Print("Masukkan jumlah daerah kerabat: ")
	fmt.Scanln(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah daerah harus di antara 1 dan 999.")
		return
}

		var daerahRumah [][]int
		for i := 0; i < n; i++ {
		fmt.Printf("Masukkan jumlah rumah dan nomor rumah untuk daerah %d: ", i+1)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		numStrings := strings.Fields(input)

		var rumah []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Masukkan hanya angka valid.")
				return
			}
			rumah = append(rumah, num)
		}
		daerahRumah = append(daerahRumah, rumah)
}
fmt.Println("\nKeluaran:")
for _, rumah := range daerahRumah {
	var ganjil, genap []int
	for _, num := range rumah {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
		}
		sort.Slice(ganjil, func(i, j int) bool { return ganjil[i] <ganjil[j] })
		sort.Slice(genap, func(i, j int) bool { return genap[i] >genap[j] })
		for _, num := range ganjil {
		fmt.Printf("%d ", num)
		}
for _, num := range genap {
		fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
} 